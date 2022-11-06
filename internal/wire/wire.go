//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	nsq "github.com/nsqio/go-nsq"

	"github.com/gopperin/sme-delay-service/internal/domain/base"
	"github.com/gopperin/sme-delay-service/internal/domain/delay"
)

// InitBaseAPI init base api wire
func InitBaseAPI() base.API {
	wire.Build(base.ProvideAPI, base.ProvideService)
	return base.API{}
}

// InitDelayAPI init base api wire
func InitDelayAPI(p *nsq.Producer, redis *redis.Client) delay.API {
	wire.Build(delay.ProvideAPI, delay.ProvideService)
	return delay.API{}
}
