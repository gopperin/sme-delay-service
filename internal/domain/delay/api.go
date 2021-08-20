package delay

import (
	"bytes"
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

	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "请求错误"})
		return
	}

	var msg types.DelayMsg
	err = json.Unmarshal(body, &msg)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "JSON参数解析错误"})
		return
	}

	a.Service.DeferredPublish("sme-delay-service", time.Duration(msg.Delay)*time.Second, body)
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "发送成功", "data": msg})
}

// Callback Callback
func (a *API) Callback(c *gin.Context) {

	body, err := c.GetRawData()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("data: %v\n", string(body))

	//把读过的字节流重新放到body
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": string(body)})
}
