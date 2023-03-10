package rabbitmq

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"sync"
	"trigger_listener_service/config"
)

type RabbitMQ struct {
	channel    *amqp.Channel
	consumers  map[string]*Consumer
	publishers map[string]*Publisher
}

type RabbitMQI interface {
	AddConsumer(name string, handler HandlerFunc) error
	RunConsumers(ctx context.Context)
	AddPublisher(name string) error
	Publish(ctx context.Context, name string, data []byte) error
}

func NewRabbitMQ(cfg config.Config, ch *amqp.Channel) (RabbitMQI, error) {
	rabbit := &RabbitMQ{
		consumers:  make(map[string]*Consumer),
		publishers: make(map[string]*Publisher),
		channel:    ch,
	}

	return rabbit, nil
}

// Run consumers
func (r RabbitMQ) RunConsumers(ctx context.Context) {
	var wg sync.WaitGroup

	for _, consumer := range r.consumers {
		wg.Add(1)
		go func(wg *sync.WaitGroup, c *Consumer) {
			defer wg.Done()
			for d := range c.message {
				if err := c.handler(ctx, d.Body); err != nil {
					panic(err)
				}
			}
		}(&wg, consumer)

	}

	wg.Add(1)
	wg.Wait()
}
