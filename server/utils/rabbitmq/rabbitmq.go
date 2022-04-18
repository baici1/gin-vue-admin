package rabbitmq

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/rabbitmq/hello_world"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"os"
)

func CreateRabbitMqFactory() *rmqHandleMsg {
	return &rmqHandleMsg{}
}

type rmqHandleMsg struct{}

var produceComsumeInfomationService = service.ServiceGroupApp.AutoCodeServiceGroup.ProduceComsumeInfomationService

//通过helloworld模式处理消息

func (r *rmqHandleMsg) StartHandleMsgByHelloWorld() {

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
		data := autocode.ProduceComsumeInfomation{}
		err = json.Unmarshal([]byte(receivedData), &data)
		if err != nil {
			global.GVA_LOG.Error("反序列化失败", zap.Error(err))
		}
		//存入mqsql里面
		err = produceComsumeInfomationService.CreateProduceComsumeInfomation(data)
		if err != nil {
			global.GVA_LOG.Error("CreateProduceComsumeInfomation failed", zap.Error(err))
		}
	})

}
