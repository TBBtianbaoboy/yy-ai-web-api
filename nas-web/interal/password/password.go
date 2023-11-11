//---------------------------------
//File Name    : password.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-17 18:36:48
//Description  :
//----------------------------------
package password

import (
	"crypto/sha256"
	"encoding/base64"
	"nas-web/web-utils"
	"strconv"
	"strings"

	"github.com/jcmturner/gofork/x/crypto/pbkdf2"
)

func MakePassword(passwd string) string {
	salt := webutils.String.GetRandomString(12)
	dk := pbkdf2.Key([]byte(passwd), []byte(salt), 36000, 32, sha256.New)
	str := base64.StdEncoding.EncodeToString(dk)
	return "nas_aico_sha256" + "$" + strconv.FormatInt(int64(36000), 10) + "$" + salt + "$" + str
}

func CheckPassword(passwd, encode string) bool {
	t := strings.SplitN(encode, "$", 4)
	algorithm := t[0]
	salt := t[2]
	iterations, _ := strconv.Atoi(t[1])
	digest := sha256.New
	// algorithm must be nas_aico_sha256
	if algorithm != "nas_aico_sha256" {
		return false
	}
	dk := pbkdf2.Key([]byte(passwd), []byte(salt), iterations, 32, digest)
	str := base64.StdEncoding.EncodeToString(dk)
	hashed := "nas_aico_sha256" + "$" + strconv.FormatInt(int64(iterations), 10) + "$" + string(salt) + "$" + str
	return hashed == encode
}
