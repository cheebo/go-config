package flags

import (
	"testing"
	"flag"
	"github.com/cheebo/go-config/flags"
	"github.com/cheebo/go-config/types"
	"github.com/stretchr/testify/assert"
)

func TestDatabaseFlags(t *testing.T) {
	assert := assert.New(t)

	var (
		url = "amqp://guest:guest@localhost/"
		exchange = "ex"
		queue = "queue1"
		kind = "fanout"
		key = "events"
		durable = true
		autoDelete = false
		deliveryMode = uint(2)
	)

	args := []string{
		"-amqp.url="+url,
		"-amqp.exchange="+exchange,
		"-amqp.queue="+queue,
		"-amqp.kind="+kind,
		"-amqp.key="+key,
		"-amqp.durable=true",
		"-amqp.auto_delete=false",
		"-amqp.delivery_mode=2",
	}

	f := flag.NewFlagSet("", flag.ContinueOnError)
	cfg := &types.AMQPConfig{}
	flags.AmqpFlags(cfg, f)

	err := f.Parse(args)

	assert.NoError(err)
	assert.Equal(url, cfg.URL)
	assert.Equal(exchange, cfg.Exchange)
	assert.Equal(queue, cfg.Queue)
	assert.Equal(kind, cfg.Kind)
	assert.Equal(key, cfg.Key)
	assert.Equal(durable, cfg.Durable)
	assert.Equal(autoDelete, cfg.AutoDelete)
	assert.Equal(deliveryMode, cfg.DeliveryMode)
}
