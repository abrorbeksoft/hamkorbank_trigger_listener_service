package trigger_listener_service

import (
	"context"
	"encoding/json"
	"fmt"
	"trigger_listener_service/pkg/logger"
)

func (t *triggerListener) ListenErrors(ctx context.Context, data []byte) error {
	var resp Message
	err := json.Unmarshal(data, &resp)
	if err != nil {
		t.log.Error("error while consuming ", logger.Error(err))
		return err
	}
	str := fmt.Sprintf("error on data %s", resp.RecordId)
	t.log.Error(str, logger.Error(err))

	return nil
}

func (t *triggerListener) ListenInfo(ctx context.Context, data []byte) error {
	var resp Message
	err := json.Unmarshal(data, &resp)
	if err != nil {
		t.log.Error("error while consuming ", logger.Error(err))
		return err
	}
	str := fmt.Sprintf("error on data %s", resp.RecordId)
	t.log.Info(str, logger.Error(err))
	return nil
}

func (t *triggerListener) ListenDebug(ctx context.Context, data []byte) error {
	var resp Message
	err := json.Unmarshal(data, &resp)
	if err != nil {
		t.log.Error("error while consuming ", logger.Error(err))
		return err
	}
	str := fmt.Sprintf("error on data %s", resp.RecordId)
	t.log.Debug(str, logger.Error(err))
	return nil
}

func (t *triggerListener) ListenAll(ctx context.Context, data []byte) error {
	var resp Message
	err := json.Unmarshal(data, &resp)
	if err != nil {
		t.log.Error("error while consuming ", logger.Error(err))
		return err
	}
	str := fmt.Sprintf("error on data %s", resp.RecordId)
	t.log.Info(str, logger.Error(err))
	return nil
}
