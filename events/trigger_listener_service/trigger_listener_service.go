package trigger_listener_service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"syscall"
	"trigger_listener_service/config"
	"trigger_listener_service/pkg/logger"
)

func (t *triggerListener) Listen(ctx context.Context, data []byte) error {
	var resp = &Message{}
	err := json.Unmarshal(data, resp)
	if err != nil {
		t.log.Error("error while consuming ", logger.Error(err))
		return err
	}

	t.log.Info("Debug", logger.Any("resp ", resp))

	d, err := json.Marshal(Message{
		RecordId: "Nima gap",
	})
	if err != nil {
		t.log.Error("Error while marshaling data", logger.Error(err))
	}

	err = t.rabbitmq.Publish(ctx, config.AllDebug, d)
	if err != nil {
		t.log.Error("Error while publishing data", logger.Error(err))
	}

	path := fmt.Sprintf("http://%s:%d/v1/phone/%s", t.cfg.RestServiceHost, t.cfg.RestServicePort, resp.RecordId)

	_, status, err := t.httpClient.Request("GET", path, "application/json", "", nil, "")
	if errors.Is(err, syscall.ECONNREFUSED) {
		panic(err)
	}

	b, err := json.Marshal(Message{
		RecordId: resp.RecordId,
	})
	if err != nil {
		t.log.Error("Error while marshaling data", logger.Error(err))
	}

	if status == 404 {
		err = t.rabbitmq.Publish(ctx, config.AllErrors, b)
		if err != nil {
			t.log.Error("Error while publishing data", logger.Error(err))
			return err
		}
		return nil
	}

	if status == 500 {
		err = t.rabbitmq.Publish(ctx, config.Consumer, b)
		if err != nil {
			t.log.Error("Error while publishing data", logger.Error(err))
			return err
		}
		return nil
	}

	err = t.rabbitmq.Publish(ctx, config.AllInfo, b)
	if err != nil {
		t.log.Error("Error while publishing data", logger.Error(err))
		return err
	}

	return nil
}
