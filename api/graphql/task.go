package graphql

import (
	"context"
	"fmt"
	"time"

	"github.com/aos-dev/go-toolbox/zapcontext"
	"go.uber.org/zap"

	"github.com/aos-dev/dm/models"
)

// runTask handle publish task and update
func (r *mutationResolver) runTask(ctx context.Context, task *models.Task) error {
	gc := GinContextFrom(ctx)
	logger := zapcontext.FromGin(gc)

	pt, err := task.FormatProtoTask()
	if err != nil {
		logger.Error("format proto task", zap.String("dm task", task.ID), zap.Error(err))
		return fmt.Errorf("format proto task failed: %w", err)
	}
	if err = r.Manager.Publish(ctx, pt); err != nil {
		logger.Error("publish task", zap.String("dm task", task.ID), zap.String("task", pt.Id), zap.Error(err))
		return err
	}

	var status models.TaskStatus
	// TODO: Exec task asynchronously
	err = r.Manager.Wait(ctx, pt)
	if err != nil {
		logger.Error("task wait", zap.String("dm task", task.ID), zap.String("task", pt.Id), zap.Error(err))
		task.Status = models.StatusError
	} else {
		logger.Info("task exec succeed", zap.String("dm task", task.ID), zap.String("task", pt.Id))
		// set task status into finished async
		task.Status = models.StatusFinished
	}
	// update time
	task.UpdatedAt = time.Now()
	err = r.DB.SaveTask(task)
	if err != nil {
		logger.Error("set task status", zap.String("dm task", task.ID),
			zap.String("status", status.String()), zap.Error(err))
	}
	return nil
}
