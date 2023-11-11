//---------------------------------
//File Name    : jwt.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-17 14:44:24
//Description  :
//----------------------------------
package jwt

import (
	"fmt"
	"nas-common/mlog"
	"nas-common/models"
	"nas-web/dao/redis"
	"nas-web/support"
	"strings"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"go.uber.org/zap"
)

type (
	errorHandler func(context.Context,string) // 用来处理错误
	TokenExtractor func(context.Context)(string,error) // 用来提取jwt token
	Jwts struct{
		Config Config
	}
)

var (
	jwts      *Jwts
	lock      sync.Mutex
	JwtSecKey = support.JWTFixedSecKey
	SysToken  string //System info 防jwt重放
)

//Func 用于校验jwt token 是否正确
func Server(ctx context.Context) bool {
	ConfigJWT()
	if err := jwts.CheckJWT(ctx); err != nil {
		mlog.Debug("Check jwt error", zap.Error(err), zap.String("url", ctx.Request().RequestURI))
		return false
	}
	return true
}

//Func 获取请求头部中的jwt token (nas)
func GetAuthHeader(ctx context.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return "", nil
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "nas" {
		return "", fmt.Errorf("Authorization header format must be nas {token} ")
	}
	return authHeaderParts[1], nil
}

//Func 配置jwt中间件
func ConfigJWT() {
	if jwts != nil {
		return
	}

	lock.Lock()
	defer lock.Unlock()

	cfg := Config{
		ContextKey: DefaultContextKey,
		ValidationKeyGetter: func(token *jwt.Token) (i interface{}, e error) {
			return []byte(JwtSecKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		ErrorHandler: func(ctx context.Context, errMsg string) {
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.StopExecution()
		},
		Extractor:           GetAuthHeader,
		Expiration:          true,
		Debug:               true,
		EnableAuthOnOptions: false,
	}
	jwts = &Jwts{Config: cfg}
}

//Func 检验jwt的token是否正确
func (j *Jwts) CheckJWT(ctx context.Context) error {
	if !j.Config.EnableAuthOnOptions {
		if ctx.Method() == iris.MethodOptions {
			return nil
		}
	}

	// 使用自定义token提取函数(FromAuthHeader)
	token, err := j.Config.Extractor(ctx)
	if err != nil {
		j.logf("Error extracting token: %v", err)
		return fmt.Errorf("Error extracting token: %v ", err)
	}

	// Jwt token 为空时
	if token == "" {
		j.logf("  Error: No credentials found (CredentialsOptional=false)")
		return fmt.Errorf(support.TokenParseFailedAndEmpty)
	}

	// 校验jwt的token是否在黑名单中
	if redis.IsJwtInBlackList(token) {
		return fmt.Errorf(support.TokenExpire)
	}
	// 校验jwt的token是否在白名单中
	if !redis.IsJwtInWhiteList(token){
		return fmt.Errorf(support.TokenExpire)
	} else {
		// jwt token正确则自动续期
		if err := redis.FlushJwtWhiteList(token); err != nil {
			return fmt.Errorf(support.TokenFlushFailed)
		}
	}
	// 进行Jwt Token 的解析
	parseToken, err := jwt.Parse(token, j.Config.ValidationKeyGetter)
	if err != nil {
		j.logf("Error parsing token: %v")
		return fmt.Errorf("Error parsing token: %v ", err)
	}
	// 数据算法校验(alg字段)
	if j.Config.SigningMethod != nil && j.Config.SigningMethod.Alg() != parseToken.Header["alg"] {
		message := fmt.Sprintf("Expected %s signing method but token specified %s",
			j.Config.SigningMethod.Alg(), parseToken.Header["alg"])
		j.logf("Error validating token algorithm: %s", message)
	}
	// 数据字段校验
	if !parseToken.Valid { //参数值为空时
		j.logf(support.TokenParseFailedAndInvalid)
		return fmt.Errorf(support.TokenParseFailedAndInvalid)
	}
	if claims, ok := parseToken.Claims.(jwt.MapClaims); ok {
		user := &models.UserToken{
			UserId:       int(claims["userId"].(float64)),
			UserType:      int(claims["userType"].(float64)),
		}
		ctx.Values().Set(j.Config.ContextKey, user)
	}
	return nil
}

// jwt 续期时间计算函数
func remainingTime(exp float64) (int64, error) {
	if int64(exp) > time.Now().Unix() {
		return int64(exp) - time.Now().Unix(), nil
	}
	return 0, fmt.Errorf("get claims expire time error, expireTime: %f", exp)
}

// jwt异常日志
func (j *Jwts) logf(format string, args ...interface{}) {
	if j.Config.Debug {
		mlog.Debugf(format, args...)
	}
}

type Claims struct {
	UserId       int `json:"userId"`
	UserType     int `json:"userType"`
	jwt.StandardClaims
}

