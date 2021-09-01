package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	nsq "github.com/nsqio/go-nsq"

	mywire "sme-delay-service/internal/wire"
)

// SetupDelayRouter SetupDelayRouter
func SetupDelayRouter(g *gin.Engine, p *nsq.Producer, redis *redis.Client) {

	// initialize API
	api := mywire.InitDelayAPI(p, redis)

	r := g.Group("/api/v1/delay")
	{
		r.POST("/send", api.Send)

		r.POST("/callback", api.Callback)
	}
}
