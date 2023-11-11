//---------------------------------
//File Name    : producer.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2022-01-09 17:08:40
//Description  :
//----------------------------------
package mkafka

import (
	"sync"

	"github.com/Shopify/sarama"
)

type ProducerOption func(*Producer) error

//用户消息 ok
type ProducerMessage struct {
	//topic
	Topic string
	// The partitioning key for this message. Pre-existing Encoders include
	// StringEncoder and ByteEncoder.
	Key sarama.Encoder
	//Payload to produce.
	Payload []byte
	// The headers are key-value pairs that are transparently passed
	// by Kafka between producers and consumers.
	Headers []sarama.RecordHeader
	//Custom data which will be enqueued FailedQueue if failed,
	//or be enqueued SucceedQueue if succeed.
	Custom interface{}
	//Partition specified by user when OptionManualPartition
	//was set, or returned by broker.
	Partition int32
	//Offset returned by broker.
	Offset int64
	//Error will be set while showing up in FailedQueue.
	Error error
}

/*
*	Producer结束应该调用Fini，这样会等待本模块所有可能block
*	的函数结束才返回。 ok
**/
type Producer struct {
	unshare bool
	queue   uint64
	concu   uint32

	failedCh  chan<- *ProducerMessage
	succeedCh chan<- *ProducerMessage
	inputCh   chan *ProducerMessage

	tps *sync.Map

	addrs  []string
	config *sarama.Config
	asyncP sarama.AsyncProducer

	wg *sync.WaitGroup
}

type topicParamsProducer struct {
	queue        uint64
	concu        uint32
	topicInputCh chan *ProducerMessage

	asyncP sarama.AsyncProducer
	w      *workerP
}

func NewProducer(addrs []string, options ...ProducerOption) (*Producer, error) {
	p := &Producer{
		queue:  0,
		concu:  1,
		addrs:  addrs,
		config: sarama.NewConfig(),
		tps:    new(sync.Map),
		wg:     new(sync.WaitGroup),
	}
	for _, option := range options {
		option(p)
	}
	p.inputCh = make(chan *ProducerMessage, p.queue)
	if !p.unshare {
		asyncP, err := sarama.NewAsyncProducer(addrs, p.config)
		if err != nil {
			return nil, err
		}
		p.asyncP = asyncP
	}

	p.tps.Range(func(key, value interface{}) bool {
		tp := value.(*topicParamsProducer)
		p.spawn(tp)
		return true
	})

	go p.dispatch()
	return p, nil
}

func (p *Producer) Input() chan<- *ProducerMessage {
	return p.inputCh
}

func (p *Producer) Fini() {
	if !p.unshare {
		p.asyncP.Close()
	}

	p.tps.Range(func(key, value interface{}) bool {
		tp := value.(*topicParamsProducer)
		if p.unshare {
			tp.asyncP.Close()
		}

		close(tp.topicInputCh)
		tp.w.fini()

		return true
	})

	close(p.inputCh)
	p.wg.Wait()
	return
}

//dispatch函数会在close(p.inputCh)后返回
func (p *Producer) dispatch() {
	p.wg.Add(int(p.concu))
	for i := 0; i < int(p.concu); i++ {
		go func() {
			defer p.wg.Done()

			for msg := range p.inputCh {
				tpIf, ok := p.tps.Load(msg.Topic)
				if !ok {
					var loaded bool
					tp := &topicParamsProducer{
						queue:        0,
						concu:        1,
						topicInputCh: make(chan *ProducerMessage, 0),
					}
					tpIf, loaded = p.tps.LoadOrStore(msg.Topic, tp)
					if !loaded {
						err := p.spawn(tp)
						if err != nil && p.failedCh != nil {
							msg.Error = err
							p.failedCh <- msg
						}
					}
				}
				tp := tpIf.(*topicParamsProducer)
				tp.topicInputCh <- msg
			}
		}()
	}
}

func (p *Producer) spawn(tp *topicParamsProducer) error {
	var asyncP sarama.AsyncProducer
	var err error
	if p.unshare {
		asyncP, err = sarama.NewAsyncProducer(p.addrs, p.config)
		if err != nil {
			return err
		}
		tp.asyncP = asyncP
	} else {
		tp.asyncP = p.asyncP
	}
	w := newworkerP()
	w.work(tp.topicInputCh, p.failedCh, p.succeedCh, tp.asyncP, tp.concu)
	tp.w = w
	return nil
}

func OptionFailedCh(ch chan<- *ProducerMessage) ProducerOption {
	return func(p *Producer) error {
		p.config.Producer.Return.Errors = true
		p.failedCh = ch
		return nil
	}
}

func OptionVersion(version sarama.KafkaVersion) ProducerOption {
	return func(p *Producer) error {
		p.config.Version = version
		return nil
	}
}

type workerP struct {
	wg *sync.WaitGroup
}

func newworkerP() *workerP {
	w := &workerP{
		wg: new(sync.WaitGroup),
	}
	return w
}

func (w *workerP) work(inCh <-chan *ProducerMessage, fCh, sCh chan<- *ProducerMessage, asyncP sarama.AsyncProducer, concu uint32) {
	w.wg.Add(int(concu))
	for i := 0; i < int(concu); i++ {
		go func() {
			defer w.wg.Done()
			for msg := range inCh {
				saramaMsg := &sarama.ProducerMessage{
					Topic:     msg.Topic,
					Key:       msg.Key,
					Value:     sarama.ByteEncoder(msg.Payload),
					Headers:   msg.Headers,
					Metadata:  msg.Custom,
					Partition: msg.Partition,
				}
				asyncP.Input() <- saramaMsg
			}
		}()
	}

	if fCh != nil {
		w.wg.Add(1)
		go func() {
			defer w.wg.Done()
			for pe := range asyncP.Errors() {
				payload, _ := pe.Msg.Value.Encode()
				msg := &ProducerMessage{
					Topic:     pe.Msg.Topic,
					Key:       pe.Msg.Key,
					Payload:   payload,
					Headers:   pe.Msg.Headers,
					Custom:    pe.Msg.Metadata,
					Partition: pe.Msg.Partition,
					Offset:    pe.Msg.Offset,
					Error:     pe.Err,
				}
				fCh <- msg
			}
		}()
	}

	if sCh != nil {
		w.wg.Add(1)
		go func() {
			defer w.wg.Done()
			for saramaMsg := range asyncP.Successes() {
				payload, _ := saramaMsg.Value.Encode()
				msg := &ProducerMessage{
					Topic:     saramaMsg.Topic,
					Key:       saramaMsg.Key,
					Payload:   payload,
					Headers:   saramaMsg.Headers,
					Custom:    saramaMsg.Metadata,
					Partition: saramaMsg.Partition,
					Offset:    saramaMsg.Offset,
				}
				sCh <- msg
			}
		}()
	}
}

func (w *workerP) fini() {
	w.wg.Wait()
}
