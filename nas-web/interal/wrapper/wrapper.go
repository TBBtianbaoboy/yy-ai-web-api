//---------------------------------
//File Name    : wrapper.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-14 17:35:36
//Description  :
//----------------------------------
package wrapper

import (
	"encoding/json"
	"nas-common/mlog"
	"nas-web/support"
	"reflect"
	"strings"

	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"
)
type ApiHandler func(ctx *Context, reqBody interface{}) error

type ApiConfig struct {
	ReqType support.CheckType
}

var validate = validator.New()

func ApiWrapper(ctx *Context, handler ApiHandler, paramChecker bool, reqBody interface{}, params ...interface{}) {
	var err error
	var paramErr map[string]string
	//请求Body校验
	if reqBody != nil {
		if len(params) == 0 { //参数为空
			mlog.Error("ApiWrapper传入的params is nil")
			support.SendApiErrorResponse(ctx, "params is empty", iris.StatusInternalServerError)
			return
		}
		//请求类型校验 form|json
		config := params[0].(*ApiConfig)
		switch config.ReqType {
		case support.CHECKTYPE_FORM:
			err = ctx.ReadForm(reqBody)
		case support.CHECKTYPE_JSON:
			err = ctx.ReadJSON(reqBody)
		}
		if err != nil && !iris.IsErrPath(err) {
			mlog.Error("Api Wrapper reqBody parse failed", zap.Error(err))
			support.SendApiErrorResponse(ctx, "parse reqBody failed", iris.StatusInternalServerError)
			return
		}
		//请求参数校验
		if paramChecker {
			if err, paramErr = checkParam(config.ReqType, reqBody); err != nil || paramErr != nil {
				if err != nil {
					mlog.Error("checker param error", zap.Error(err))
					support.SendApiErrorResponse(ctx, "checker param failed failed", iris.StatusOK)
				} else {
					if msg, err := json.Marshal(paramErr); err == nil {
						support.SendApiErrorResponse(ctx, string(msg), iris.StatusOK)
					} else {
						mlog.Error("marshal err param msg failed", zap.Error(err))
						support.SendApiErrorResponse(ctx, "checker param failed failed", iris.StatusInternalServerError)
					}
				}
				return
			}
		}
	}
	if err := handler(ctx, reqBody); err != nil {
		mlog.Error("handler exec failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, "handler exec failed", iris.StatusInternalServerError)
	}
}

// 参数校验函数
func checkParam(reqType support.CheckType, reqBody interface{}) (error, map[string]string) {
	if err := validate.Struct(reqBody); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err, nil
		}
		paramErr := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			st := reflect.TypeOf(reqBody)
			if param, ok := st.Elem().FieldByName(strings.Split(err.StructField(), "[")[0]); ok {
				if reqType == support.CHECKTYPE_JSON {
					tag := param.Tag.Get("json")
					paramErr[tag] = "Invalid input"
				} else if reqType == support.CHECKTYPE_FORM {
					tag := param.Tag.Get("form")
					paramErr[tag] = "Invalid input"
				}
			}
		}
		return nil, paramErr
	}
	return nil, nil
}
