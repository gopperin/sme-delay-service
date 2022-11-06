package api

import (
	"log"
	"time"

	nsq "github.com/nsqio/go-nsq"

	myconfig "github.com/gopperin/sme-delay-service/internal/config"
	"github.com/gopperin/sme-delay-service/internal/domain/delay"
)

var nsqConfig *nsq.Config

func init() {
	nsqConfig = nsq.NewConfig()
	nsqConfig.DefaultRequeueDelay = 0
	nsqConfig.MaxBackoffDuration = 20 * time.Millisecond
	nsqConfig.LookupdPollInterval = 1000 * time.Millisecond
	nsqConfig.RDYRedistributeInterval = 1000 * time.Millisecond
	nsqConfig.MaxInFlight = 2500
}

// HandleDelayMessage HandleDelayMessage
func HandleDelayMessage(message *nsq.Message) error {
	go delay.DoneDelay(string(message.Body))
	return nil
}

// MakeConsumer MakeConsumer
func MakeConsumer(topic, channel string, config *nsq.Config,
	handle func(message *nsq.Message) error) {
	consumer, _ := nsq.NewConsumer(topic, channel, config)
	consumer.AddHandler(nsq.HandlerFunc(handle))

	// 待深入了解
	// 連線到 NSQ 叢集，而不是單個 NSQ，這樣更安全與可靠。
	// err := q.ConnectToNSQLookupd("127.0.0.1:4161")

	err := consumer.ConnectToNSQD(myconfig.Case.Nsq.Addr)
	if err != nil {
		log.Panic("Could not connect")
	}
}
