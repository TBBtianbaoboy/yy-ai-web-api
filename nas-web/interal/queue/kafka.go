//---------------------------------
//File Name    : kafka.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2022-01-10 13:29:33
//Description  :
//----------------------------------
package queue

import (
	"fmt"
	"nas-web/config"
	"nas-web/pkg/dispatcher"

	"nas-common/msingleton"
)

var (
	dispatcherKafkaMgr msingleton.Singleton
)

func DispatcherKafkaInst() dispatcher.Dispatcher {
	mgr, err := dispatcherKafkaMgr.Get()
	if err != nil {
		return nil
	}
	return mgr.(dispatcher.Dispatcher)
}

func KafkaInit() error {
	dispatcherKafkaMgr = msingleton.NewSingleton(func() (interface{}, error) {
		var d dispatcher.Dispatcher

		d, err := dispatcher.New(
			dispatcher.WithType(dispatcher.TypeKafka),
			dispatcher.WithKafkaEndpoints(config.IrisConfig.Kafka.EndPoints...),
		)
		if err != nil {
			return nil, err
		}

		return d, nil
	})

	if DispatcherKafkaInst() == nil {
		return fmt.Errorf("DispatcherKafkaInst is nil")
	}

	return nil
}
