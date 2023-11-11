//---------------------------------
//File Name    : captcha.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-14 18:36:54
//Description  : 
//----------------------------------
package cache

import (
	"github.com/mojocn/base64Captcha"
)


var captchaCharacterOpt = base64Captcha.ConfigCharacter{
	Height:             80,
	Width:              240,
	Mode:               3,
	ComplexOfNoiseText: 0,
	ComplexOfNoiseDot:  0,
	IsUseSimpleFont:    false,
	IsShowHollowLine:   false,
	IsShowNoiseDot:     false,
	IsShowNoiseText:    false,
	IsShowSlimeLine:    false,
	IsShowSineLine:     false,
	CaptchaLen:         4,
}

// 生成验证码模块.
func GenDigitCaptcha() (id, pngData string) {
	captchaId, characterCap := base64Captcha.GenerateCaptcha("", captchaCharacterOpt)
	base64Png := base64Captcha.CaptchaWriteToBase64Encoding(characterCap)
	return captchaId, base64Png
}

// 校验验证码是否正确
func VerifyCaptcha(id, value string) bool {
	return base64Captcha.VerifyCaptcha(id, value)
}

