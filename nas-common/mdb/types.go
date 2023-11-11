//---------------------------------
//File Name    : types.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-08 11:12:30
//Description  :
//----------------------------------
package mdb

import (
	"github.com/globalsign/mgo"
	"github.com/pkg/errors"
)

var (
	ErrorLimit      = errors.New("find limit is invalid,must be -1 or > 0")
	ErrorResultType = errors.New("error result type")
	ErrNotFound     = mgo.ErrNotFound
)

// 数据库操作接口封装. 要支持mongodb实例,mysql等.
type DBAdaptor interface {
	Connect(interface{}) error
	SetPoolLimit(limit int)

	Close()

	// 获取原始指针.
	GetOrg() interface{}

	// 常用操作接口.
	FindOne(name string, query, result interface{}) (exist bool, err error)
	Find(name string, query, result interface{}, limit int) error
	FindAll(name string, query, result interface{}) error
	FindAllWithSelect(name string, query, selection, result interface{}) error
	FindOneWithSelect(name string, query, selection, result interface{}) (exist bool, err error)

	RemoveAll(name string, query interface{}) error

	Insert(name string, docs ...interface{}) error
	Update(name string, query interface{}, update interface{}, multi bool) error
	UpdateManual(name string, query interface{}, update interface{}, multi bool) error

	FindWithSelect(name string, query, selection, result interface{}, limit int) error
	FindCount(name string, query interface{}) (c int, err error)
	FindByLimitAndSkip(name string, query, result interface{}, limit, skip int) error
	FindSortByLimitAndSkip(name string, query interface{}, result interface{}, limit, skip int, sorter ...string) error
}
