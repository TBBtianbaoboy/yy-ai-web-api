//---------------------------------
//File Name    : token.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-14 17:16:30
//Description  :
//----------------------------------
package models

//@Struct 存储用户token 信息
type UserToken struct {
	UserId   int    `json:"user_id" bson:"user_id"`     //用户Id
	UserType int `json:"user_type" bson:"user_type"` //用户类型
}
