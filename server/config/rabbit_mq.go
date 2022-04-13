package config

type Rabbitmq struct {
	HelloWorld HelloWorld `yaml:"helloWorld"`
}

type HelloWorld struct {
	ConsumerChanNumber          int    `yaml:"consumerChanNumber"`
	OffLineReconnectIntervalSec int    `yaml:"offLineReconnectIntervalSec"`
	RetryCount                  int    `yaml:"retryCount"`
	Addr                        string `yaml:"addr"`
	QueueName                   string `yaml:"queueName"`
	Durable                     bool   `yaml:"durable"`
}
