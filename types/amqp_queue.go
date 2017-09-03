package types

type AmqpQueue struct {
	AMQPUrl      string `json:"amqp_url"`
	ExchangeName string `json:"exchange_name"`
	QueueName    string `json:"queue_name"`
	Key          string `json:"key"`
	Kind         string `json:"kind"`
	AutoDelete   bool   `json:"auto_delete"`
	Durable      bool   `json:"durable"`
}