package router

import (
	"github.com/gin-gonic/gin"

	mywire "sme-delay-service/internal/wire"
)

// SetupBaseRouter SetupBaseRouter
func SetupBaseRouter(g *gin.Engine) {

	// initialize API
	api := mywire.InitBaseAPI()

	r := g.Group("/")
	{
		r.GET("health", api.Health)
		r.GET("release", api.Release)
	}
}
