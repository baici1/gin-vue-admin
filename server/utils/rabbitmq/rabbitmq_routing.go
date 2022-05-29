package rabbitmq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/rabbitmq/routing"
	"github.com/flipped-aurora/gva-plugins/email/service"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"os"
	"text/template"
)

func CreateRabbitMqFactoryByRouting() *rmqHandleMsgByRouting {
	return &rmqHandleMsgByRouting{}
}

type rmqHandleMsgByRouting struct{}

func (r *rmqHandleMsgByRouting) StartHandleMsgByRoutingToStudent() {
	consumer, err := routing.CreateConsumer()
	if err != nil {
		global.GVA_LOG.Error("Routing create consumer failed", zap.Error(err))
		os.Exit(1)
	}
	consumer.OnConnectionError(func(err *amqp.Error) {
		global.GVA_LOG.Error("Routing Connection error:", zap.Error(err))
	})
	consumer.Received("student", func(receivedData string) {
		fmt.Printf("student回调函数处理消息：--->%s\n", receivedData)
		//反序列化
		data := autocodeReq.RecruitToRabbitmqInfo{}
		err = json.Unmarshal([]byte(receivedData), &data)
		//解析邮件模板文件
		tmpl, err := template.ParseFiles(utils.ConfigTemplate)
		if err != nil {
			global.GVA_LOG.Error("文件解析失败", zap.Error(err))
			return
		}
		//存储body
		var body bytes.Buffer
		body.Write([]byte{})
		err = tmpl.Execute(&body, data)
		if err != nil {
			global.GVA_LOG.Error("模板写入失败", zap.Error(err))
			return
		}
		fmt.Println(data.Comsumer.Email)
		err = service.ServiceGroupApp.SendEmail(data.Comsumer.Email, "恭喜你有人接受了你的要求！", body.String())
		//err = utils.Email("1308964967@qq.com", "恭喜你有人接受了你的要求！", body.String())
		if err != nil {
			global.GVA_LOG.Error("邮件发送失败", zap.Error(err))
			return
		}
		//开始发邮件
		if err != nil {
			global.GVA_LOG.Error("反序列化失败", zap.Error(err))
		}
	})
}
