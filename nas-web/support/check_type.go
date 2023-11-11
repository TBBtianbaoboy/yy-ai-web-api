//---------------------------------
//File Name    : check_type.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-14 17:35:04
//Description  : 
//----------------------------------
package support

type CheckType = int32

const (
	_ CheckType = iota
	CHECKTYPE_FORM
	CHECKTYPE_JSON
)
