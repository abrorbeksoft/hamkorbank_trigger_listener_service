package main

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"trigger_listener_service/config"
	"trigger_listener_service/events"
	"trigger_listener_service/pkg/logger"
	"trigger_listener_service/pkg/requests"
)

func main() {

	cfg := config.Load()
	log := logger.NewLogger(cfg.LogLevel, "trigger_listener_service")

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", cfg.RabbitMQUser, cfg.RabbitMQPassword, cfg.RabbitMQHost, cfg.RabbitMQPort))
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()

	httpClient := requests.NewHttpClient("", 10)

	pubSubServer, err := events.NewEvents(cfg, log, ch)
	if err != nil {
		log.Error("error on event server")
	}

	ctx := context.Background()
	pubSubServer.InitServices(ctx, cfg, httpClient) // it should run forever if there is any consumer
}
