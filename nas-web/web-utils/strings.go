//---------------------------------
//File Name    : web-utils/strings.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-18 17:12:00
//Description  :
//----------------------------------
package webutils

import (
	randomId "github.com/satori/go.uuid"
	"nas-common/utils"
	"regexp"
	"strings"
)

type str struct{}

var String str

func (str) Compare(str1, str2 string) bool {
	if strings.Compare(str1, str2) == 0 {
		return true
	}
	return false
}

func (str) GetPasswordStrength(passwd string) (passwdLevel int) {
	counter := 0
	numberMatch := "[0-9]+"
	lowLetter := "[a-z]+"
	upLetter := "[A-Z]+"
	specialSymbol := `[*()~!@#$%^&*-+=_|:;'<>,.?/\[\]\{\}<>]+`
	if match, _ := regexp.MatchString(numberMatch, passwd); match {
		counter++
	}
	if match, _ := regexp.MatchString(lowLetter, passwd); match {
		counter++
	}
	if match, _ := regexp.MatchString(upLetter, passwd); match {
		counter++
	}
	if match, _ := regexp.MatchString(specialSymbol, passwd); match {
		counter++
	}
	if len(passwd) < 8 || counter <= 1 {
		passwdLevel = 0
	} else if counter <= 2 {
		passwdLevel = 1
	} else if counter <= 3 {
		passwdLevel = 2
	} else {
		passwdLevel = 3
	}
	return
}

func (str) GetRandomString(len int) string {
	return utils.MD5String(randomId.NewV1().String(), len)
}

func (str) Hash(name ...string) string {
	return utils.MD5String(strings.Join(name, ":"), 32)
}
