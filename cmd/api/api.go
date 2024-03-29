package api

import (
	"net/http"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	myconfig "github.com/gopperin/sme-delay-service/internal/config"
	myrouter "github.com/gopperin/sme-delay-service/internal/router"
)

// StartCmd api
var (
	StartCmd = &cobra.Command{
		Use:   "start",
		Short: "start sme-delay-service api", SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		Run: func(cmd *cobra.Command, args []string) {
			//启动API服务
			run()

			logrus.Println("sme-delay-service end")
		},
	}
)

func setup() {

	//1. 读取配置
	myconfig.Setup("./")

}

func run() {

	router := gin.Default()

	MakeConsumer("sme-delay-service", "ch", nsqConfig, HandleDelayMessage)

	producer, err := InitProducer()
	if err != nil {
		logrus.Println("InitProducer error", err.Error())
		return
	}
	logrus.Println("InitProducer success")

	redis, err := InitRedis()
	if err != nil {
		logrus.Println("InitRedis error", err.Error())
		return
	}
	logrus.Println("InitRedis success")

	logrus.Info("sme-delay-service start on:", myconfig.Case.Application.Port)

	/* api base */
	myrouter.SetupBaseRouter(router)

	/* product base */
	myrouter.SetupDelayRouter(router, producer, redis)

	server := &http.Server{
		Addr:         ":" + myconfig.Case.Application.Port,
		Handler:      router,
		ReadTimeout:  300 * time.Second,
		WriteTimeout: 300 * time.Second,
	}

	logrus.Println("sme-delay-service start on:", myconfig.Case.Application.Port)
	gracehttp.Serve(server)
}
