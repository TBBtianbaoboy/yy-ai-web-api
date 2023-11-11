//---------------------------------
//File Name    : dispatcher.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2022-01-10 13:34:20
//Description  : 
//----------------------------------
package dispatcher

import (
	"errors"
	"fmt"
	"sync"

	"nas-common/mkafka"
	"nas-common/mlog"
	"github.com/Shopify/sarama"
)

type Type int

const (
	_ Type = iota
	TypeError
	TypeKafka
)

type Dispatcher interface {
	SetContext(data interface{}) Dispatcher // 不同的 dispatcher 设置不同的参数.
	Close()
	Dispatch(data []byte) ([]byte, error)
}

type Config struct {
	dispatcherType Type

	kafkaHost     []string
	kafkaFail     chan *mkafka.ProducerMessage
	kafkaProducer *mkafka.Producer

	once sync.Once
}

func (c *Config) Close() {
	c.once.Do(func() {
		if c.kafkaProducer != nil {
			c.kafkaProducer.Fini()
		}
	})
}

func (c *Config) Init() error {
	switch c.dispatcherType {
	case TypeKafka:
		if len(c.kafkaHost) == 0 {
			return errors.New("without kafka endpoints")
		}
		if c.kafkaFail == nil {
			c.kafkaFail = make(chan *mkafka.ProducerMessage, 100)
		}
		var err error
		c.kafkaProducer, err = mkafka.NewProducer(
			c.kafkaHost,
			mkafka.OptionFailedCh(c.kafkaFail),
			mkafka.OptionVersion(sarama.V2_3_0_0),
		)
		if err != nil {
			return err
		}

		go func() {
			for msg := range c.kafkaFail {
				if msg.Error != nil {
					mlog.Errorf("send msg to kafka fail, err:%s", msg.Error.Error())
				}
			}
		}()
	}

	return nil
}

type ConfigHandler func(c *Config)

func WithType(t Type) ConfigHandler {
	return func(c *Config) {
		c.dispatcherType = t
	}
}

func WithKafkaEndpoints(addrs ...string) ConfigHandler {
	return func(c *Config) {
		c.kafkaHost = addrs
	}
}

func New(chs ...ConfigHandler) (Dispatcher, error) {
	cfg := &Config{}

	for _, ch := range chs {
		ch(cfg)
	}

	err := cfg.Init()
	if err != nil {
		return nil, err
	}

	switch cfg.dispatcherType {
	case TypeKafka:
		return &KafkaDispatcher{Config: cfg}, nil
	default:
		return nil, fmt.Errorf("unsupported dispatcher type")
	}
}
