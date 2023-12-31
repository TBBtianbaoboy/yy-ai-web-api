//---------------------------------
//File Name    : msingleton.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2022-01-08 15:11:50
//Description  : 单例模式通用接口
//----------------------------------
package msingleton

import (
	"sync"
	"sync/atomic"
)

type SingletonInitFunc func() (interface{}, error)

type Singleton interface {
	// Return the encapsulated singleton
	Get() (interface{}, error)
}

// Call to create a new singleton that is instantiated with the given init function.
// init is not called until the first invocation of Get().  If init errors, it will be called again
// on the next invocation of Get().
func NewSingleton(init SingletonInitFunc) Singleton {
	return &singletonImpl{init: init}
}

type singletonImpl struct {
	sync.Mutex

	// The actual singleton object
	data interface{}
	// Constructor for the singleton object
	init SingletonInitFunc
	// Non-zero if init was run without error
	initialized int32
}

func (s *singletonImpl) Get() (interface{}, error) {
	// Don't lock in the common case
	if atomic.LoadInt32(&s.initialized) > 0 {
		return s.data, nil
	}

	s.Lock()
	defer s.Unlock()

	if atomic.LoadInt32(&s.initialized) > 0 {
		return s.data, nil
	}

	var err error
	s.data, err = s.init()
	if err != nil {
		return nil, err
	}

	atomic.StoreInt32(&s.initialized, 1)
	return s.data, nil
}
