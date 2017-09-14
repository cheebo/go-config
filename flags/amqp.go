package flags

import (
	"flag"
	"github.com/cheebo/go-config/types"
)

func AmqpFlags(cfg *types.AMQPConfig, f *flag.FlagSet) {
	f.StringVar(&cfg.URL, "amqp.url", cfg.URL, "AMQP url, default: amqp://guest:guest@localhost/")
	f.StringVar(&cfg.Exchange, "amqp.exchange", cfg.Exchange, "Exchange name")
	f.StringVar(&cfg.Queue, "amqp.queue", cfg.Queue, "Queue name")
	f.StringVar(&cfg.Kind, "amqp.kind", cfg.Kind, "Kind")
	f.StringVar(&cfg.Key, "amqp.key", cfg.Key, "Routing key")
	f.BoolVar(&cfg.Durable, "amqp.durable", cfg.Durable, "Durable, default: false")
	f.BoolVar(&cfg.AutoDelete, "amqp.auto_delete", cfg.AutoDelete, "Auto delete, default: true")
	f.UintVar(&cfg.DeliveryMode, "amqp.delivery_mode", cfg.DeliveryMode, "Delivery mode, default: 0")
}
