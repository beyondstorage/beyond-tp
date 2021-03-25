package graphql

import (
	"context"

	"github.com/aos-dev/go-toolbox/zapcontext"
	"github.com/aos-dev/noah/proto"
	"go.uber.org/zap"

	"github.com/aos-dev/dm/models"
)

// runTask handle publish task and update
func (r *mutationResolver) runTask(ctx context.Context, taskID string, pt *proto.Task) error {
	gc := GinContextFrom(ctx)
	logger := zapcontext.FromGin(gc)

	if err := r.Manager.Publish(ctx, pt); err != nil {
		logger.Error("publish task", zap.String("dm task", taskID), zap.String("task", pt.Id), zap.Error(err))
		return err
	}

	go func() {
		var status models.TaskStatus
		err := r.Manager.Wait(ctx, pt)
		if err != nil {
			logger.Error("task wait", zap.String("dm task", taskID), zap.String("task", pt.Id), zap.Error(err))
			status = models.StatusStopped
		} else {
			logger.Info("task exec succeed", zap.String("dm task", taskID), zap.String("task", pt.Id))
			// set task status into finished async
			status = models.StatusFinished
		}

		err = r.DB.SetTaskStatus(taskID, status)
		if err != nil {
			logger.Error("set task status", zap.String("dm task", taskID), zap.String("task", pt.Id),
				zap.String("status", status.String()), zap.Error(err))
		}
	}()
	return nil
}
