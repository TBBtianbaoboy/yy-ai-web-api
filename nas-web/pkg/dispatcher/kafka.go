//---------------------------------
//File Name    : kafka.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2022-01-10 13:35:21
//Description  : 
//----------------------------------
package dispatcher

import (
	"context"
	"errors"

	"nas-common/mkafka"
	"github.com/Shopify/sarama"
)

type KafkaContext struct {
	Ctx     context.Context
	Topic   string                // 类似于redis channel
	Key     string                // partation choice "default random" ""
	Headers []sarama.RecordHeader //
}

type KafkaDispatcher struct {
	ctx *KafkaContext
	*Config
}

func (k *KafkaDispatcher) SetContext(data interface{}) Dispatcher {
	dispatcher := &KafkaDispatcher{
		Config: k.Config,
	}

	d, ok := data.(*KafkaContext)
	if ok {
		dispatcher.ctx = d
	}

	return dispatcher
}

func (k *KafkaDispatcher) Close() { k.Config.Close() }

func (k *KafkaDispatcher) Dispatch(data []byte) ([]byte, error) {
	if k.ctx == nil {
		return nil, errors.New("context is not set")
	}

	if len(k.ctx.Topic) == 0 {
		return nil, errors.New("topic is not set")
	}

	msg := &mkafka.ProducerMessage{
		Topic:   k.ctx.Topic,
		Payload: data,
		Headers: k.ctx.Headers,
	}

	if len(k.ctx.Key) > 0 {
		msg.Key = sarama.StringEncoder(k.ctx.Key)
	}

	k.kafkaProducer.Input() <- msg

	return []byte("ok"), nil
}
