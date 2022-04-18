package config

type Rabbitmq struct {
	HelloWorld HelloWorld `yaml:"helloWorld"`
	Routing    Routing    `yaml:"routing"`
}

type HelloWorld struct {
	ConsumerChanNumber          int    `yaml:"consumerChanNumber"`
	OffLineReconnectIntervalSec int    `yaml:"offLineReconnectIntervalSec"`
	RetryCount                  int    `yaml:"retryCount"`
	Addr                        string `yaml:"addr"`
	QueueName                   string `yaml:"queueName"`
	Durable                     bool   `yaml:"durable"`
}

type Routing struct {
	DelayedExchangeName         string `yaml:"DelayedExchangeName"`
	Durable                     bool   `yaml:"Durable"`
	QueueName                   string `yaml:"QueueName"`
	OffLineReconnectIntervalSec int    `yaml:"OffLineReconnectIntervalSec"`
	RetryCount                  int    `yaml:"RetryCount"`
	Addr                        string `yaml:"Addr"`
	ExchangeType                string `yaml:"ExchangeType"`
	ExchangeName                string `yaml:"ExchangeName"`
}
