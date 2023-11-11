//---------------------------------
//File Name    : middleware/jwt/token.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-17 16:59:55
//Description  :
//----------------------------------
package jwt

import (
	"nas-common/models"
	"nas-web/config"
	webutils "nas-web/web-utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12/context"
)

//Func 生成jwt token
func GenerateToken(user *models.UserToken, secretKey string, keep bool) (token string, err error) {
	var expireTime time.Time
	if keep {
		expireTime = time.Now().AddDate(1000, 0, 0)
	} else {
		expireTime = time.Now().Add(time.Duration(config.IrisConfig.Other.JwtTimeOut) * time.Second)
	}
	claims := Claims{
		user.UserId,
		user.UserType,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "nas-aico-jwt",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenClaims.SignedString([]byte(secretKey))
	return token, err
}

// 获取用户Token
func GetUserToken(ctx context.Context) *models.UserToken {
	if ctx.Values().Get(DefaultContextKey) != nil {
		return ctx.Values().Get(DefaultContextKey).(*models.UserToken)
	}
	return nil
}

func GetTokenRemainingTime(token string) int32 {
	parsedToken, err := jwt.Parse(token, jwts.Config.ValidationKeyGetter)
	if err != nil {
		jwts.logf("Error parsing token1: %v", err)
		return 0
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
		remain, err := remainingTime(claims["exp"].(float64))
		if err != nil {
			jwts.logf("Error get remain time, claims: %v", claims)
			return 0
		}
		return int32(remain)
	}
	return 0
}

func InitSysToken() {
		SysToken = webutils.String.Hash(webutils.System.GetWebIP(),webutils.System.GetWebMacAddress())
		JwtSecKey += SysToken
}
