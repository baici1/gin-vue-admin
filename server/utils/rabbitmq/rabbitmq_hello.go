package rabbitmq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/rabbitmq/hello_world"
	"github.com/flipped-aurora/gva-plugins/email/service"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"os"
	"text/template"
)

func CreateRabbitMqFactoryByHello() *rmqHandleMsgByHello {
	return &rmqHandleMsgByHello{}
}

type rmqHandleMsgByHello struct{}

//通过helloworld模式处理消息

func (r *rmqHandleMsgByHello) StartHandleMsgByHelloWorld() {

	consumer, err := hello_world.CreateConsumer()
	if err != nil {
		global.GVA_LOG.Error("HelloWorld create consumer failed", zap.Error(err))
		os.Exit(1)
	}
	consumer.OnConnectionError(func(err *amqp.Error) {
		global.GVA_LOG.Error("HelloWorld Connection error:", zap.Error(err))
	})
	consumer.Received(func(receivedData string) {
		fmt.Printf("HelloWorld回调函数处理消息：--->%s\n", receivedData)
		//反序列化
		data := autocodeReq.StudentRecruitToRabbitmq{}
		err = json.Unmarshal([]byte(receivedData), &data)
		*data.Comsumer.Num -= 1
		err = (&autocode.StudentRecruitService{}).UpdateStudentRecruit(data.Comsumer)
		if err != nil {
			global.GVA_LOG.Error("人数修改出现错误", zap.Error(err))
			return
		}
		//解析邮件模板文件
		tmpl, err := template.ParseFiles("C:\\Users\\Y\\Desktop\\gva\\server\\utils\\rabbitmq\\index.html")
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
		fmt.Println(data.Comsumer.Member.Email)
		err = service.ServiceGroupApp.SendEmail(data.Comsumer.Member.Email, "恭喜你有人接受了你的要求！", body.String())
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
