package api

import (
	nsq "github.com/nsqio/go-nsq"

	myconfig "github.com/gopperin/sme-delay-service/internal/config"
)

// InitProducer InitProducer
func InitProducer() (*nsq.Producer, error) {

	config := nsq.NewConfig()
	producer, _ := nsq.NewProducer(myconfig.Case.Nsq.Addr, config)

	return producer, nil
}
