package hello_world

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/rabbitmq/error_record"
	"github.com/streadway/amqp"
)

// CreateProducer 创建一个生产者
func CreateProducer() (*producer, error) {
	// 获取配置信息
	conn, err := amqp.Dial(global.GVA_CONFIG.Rabbitmq.HelloWorld.Addr)
	queueName := global.GVA_CONFIG.Rabbitmq.HelloWorld.QueueName
	dura := global.GVA_CONFIG.Rabbitmq.HelloWorld.Durable

	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}

	prod := &producer{
		connect:   conn,
		queueName: queueName,
		durable:   dura,
	}
	return prod, nil
}

//  定义一个消息队列结构体：helloworld 模型
type producer struct {
	connect    *amqp.Connection
	queueName  string
	durable    bool
	occurError error
}

func (p *producer) Send(data []byte) bool {

	// 获取一个通道
	ch, err := p.connect.Channel()
	p.occurError = error_record.ErrorDeal(err)

	defer func() {
		_ = ch.Close()
	}()

	// 声明消息队列
	_, err = ch.QueueDeclare(
		p.queueName, // 队列名称
		p.durable,   //是否持久化，false模式数据全部处于内存，true会保存在erlang自带数据库，但是影响速度
		!p.durable,  //生产者、消费者全部断开时是否删除队列。一般来说，数据需要持久化，就不删除；非持久化，就删除
		false,       //是否私有队列，false标识允许多个 consumer 向该队列投递消息，true 表示独占
		false,       // 队列如果已经在服务器声明，设置为 true ，否则设置为 false；
		nil,         // 相关参数
	)
	p.occurError = error_record.ErrorDeal(err)

	// 如果队列的声明是持久化的，那么消息也设置为持久化
	msgPersistent := amqp.Transient
	if p.durable {
		msgPersistent = amqp.Persistent
	}
	// 投递消息
	if err == nil {
		err = ch.Publish(
			"",          // helloworld 、workqueue 模式设置为空字符串，表示使用默认交换机
			p.queueName, //  routing key，注意：简单模式与队列名称相同
			false,
			false,
			amqp.Publishing{
				DeliveryMode: msgPersistent, //消息是否持久化，这里与保持保持一致即可
				ContentType:  "text/plain",
				Body:         data,
			})
	}
	p.occurError = error_record.ErrorDeal(err)
	if p.occurError != nil { //  发生错误，返回 false
		return false
	} else {
		return true
	}
}

//Close 发送完毕手动关闭，这样不影响send多次发送数据
func (p *producer) Close() {
	_ = p.connect.Close()
}
