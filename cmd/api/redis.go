package api

import (
	"context"

	"github.com/go-redis/redis/v8"

	myconfig "github.com/gopperin/sme-delay-service/internal/config"
)

var ctx = context.Background()

// InitRedis InitRedis
func InitRedis() (*redis.Client, error) {

	_rdb := redis.NewClient(&redis.Options{
		Addr:     myconfig.Case.Redis.Addr,
		Password: myconfig.Case.Redis.Password, // no password set
		DB:       myconfig.Case.Redis.DB,       // use default DB
	})

	_, err := _rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return _rdb, nil
}
