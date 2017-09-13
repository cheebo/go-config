package flags

import (
	"flag"
	"github.com/cheebo/go-config/types"
)

func AmqpFlags(cfg *types.AMQPConfig, f *flag.FlagSet) {
	f.StringVar(&cfg.URL, "amqp.url", "amqp://guest:guest@localhost/", "AMQP url, default: amqp://guest:guest@localhost/")
	f.StringVar(&cfg.Exchange, "amqp.exchange", "", "Exchange name")
	f.StringVar(&cfg.Queue, "amqp.queue", "", "Queue name")
	f.StringVar(&cfg.Kind, "amqp.kind", "", "Kind")
	f.StringVar(&cfg.Key, "amqp.key", "", "Routing key")
	f.BoolVar(&cfg.Durable, "amqp.durable", false, "Durable, default: false")
	f.BoolVar(&cfg.AutoDelete, "amqp.auto_delete", true, "Auto delete, default: true")
	f.UintVar(&cfg.DeliveryMode, "amqp.delivery_mode", 0, "Delivery mode, default: 0")
}
