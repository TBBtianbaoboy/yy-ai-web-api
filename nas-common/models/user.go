//---------------------------------
//File Name    : user.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-17 18:03:07
//Description  :
//----------------------------------
package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

type UserType = int

// don't use directly
const (
	_               UserType = iota
	UserType_Admin           //1 管理员
	UserType_Super           //2 超级员
	UserType_Common          //3 普通用户
)

type User struct {
	ID               bson.ObjectId `bson:"_id,omitempty"`
	UID              int           `bson:"uid" mgo:"index:2"`  //用户ID
	Enable           bool          `bson:"enable"`             //是否允许登录 (true)-允许登录、(false)-拒绝登录
	UserType         int           `bson:"user_type"`          //用户类型
	UserName         string        `bson:"username"`           //用户名
	Password         string        `bson:"password"`           //用户密码
	PasswordStrength int           `bson:"password_strength"`  //用户密码强度 1-弱、2-中、3-强
	Mail             string        `bson:"mail"`               //用户邮箱
	Mobile           string        `bson:"mobile"`             //用户手机号码
	LastPwdChangeTm  time.Time     `bson:"last_pwd_change_tm"` //最近一次修改密码时间
	LastLoginTm      time.Time     `bson:"last_login_tm"`      //最近登录时间
	InsertTm         time.Time     `bson:"insert_tm"`          //入库时间
	UpdateTm         time.Time     `bson:"update_tm"`          //更新时间
}

func (u *User) Collection() string {
	return "tb_user"
}

type UserLoginHistory struct {
	Uid           int      `bson:"uid"`            //用户Id
	UserType      UserType `bson:"user_type"`      //用户类型
	Username      string   `bson:"username"`       //用户姓名
	RemoteAddress string   `bson:"remote_address"` //登陆地址
	LoginTime     int64    `bson:"login_time"`     //登陆时间
	Enable        bool     `bson:"enable"`         //是否允许登陆
	LoginStatu    bool     `bson:"login_statu"`    //登陆状态
}

func (u *UserLoginHistory) Collection() string {
	return "tb_user_login_history"
}
