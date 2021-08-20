package delay

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	nsq "github.com/nsqio/go-nsq"
	logrus "github.com/sirupsen/logrus"

	"sme-delay-service/internal/domain/types"
)

// Service Service
type Service struct {
	Producer *nsq.Producer
}

// ProvideService ProvideService
func ProvideService(p *nsq.Producer) Service {
	return Service{Producer: p}
}

// Publish Publish
func (s *Service) Publish(topic string, body []byte) error {
	return s.Producer.Publish(topic, body)
}

// DeferredPublish DeferredPublish
func (s *Service) DeferredPublish(topic string, delay time.Duration, body []byte) error {

	formatter := &logrus.JSONFormatter{
		// 定义时间戳格式
		TimestampFormat: "2006-01-02 15:04:05",
	}
	logrus.SetFormatter(formatter)

	logrus.Println("start send ", string(body))

	return s.Producer.DeferredPublish(topic, delay, body)
}

// DoneDelay DoneDelay
func DoneDelay(msg string) error {

	formatter := &logrus.JSONFormatter{
		// 定义时间戳格式
		TimestampFormat: "2006-01-02 15:04:05",
	}
	logrus.SetFormatter(formatter)

	var delaymsg types.DelayMsg
	err := json.Unmarshal([]byte(msg), &delaymsg)
	if err != nil {
		logrus.Println("消息解析错误")
		return err
	}

	url := delaymsg.Callback

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(msg)))
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Println(err.Error())
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	logrus.Println(string(body))

	return nil
}
