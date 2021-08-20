package delay

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"sme-delay-service/internal/domain/types"
)

// API API
type API struct {
	Service Service
}

// ProvideAPI ProvideAPI
func ProvideAPI(service Service) API {
	return API{Service: service}
}

// Send Send
func (a *API) Send(c *gin.Context) {
	var msg types.DelayMsg
	err := c.BindJSON(&msg)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	body, err := json.Marshal(msg)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	a.Service.DeferredPublish("sme-delay-service", time.Duration(msg.Delay)*time.Second, body)
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "发送成功", "data": msg})
}

// Callback Callback
func (a *API) Callback(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	if body != nil {
		fmt.Println("请求body内容为:", body)
	}

	fmt.Println("callback", string(body))

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": string(body)})
}
