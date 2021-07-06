package graphql

import (
	"context"

	"github.com/beyondstorage/go-toolbox/zapcontext"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/beyondstorage/beyond-tp/models"
)

// runTask handle publish task and update
func (r *mutationResolver) runTask(ctx context.Context, task *models.Task) error {
	gc := GinContextFrom(ctx)
	logger := zapcontext.FromGin(gc)

	task.UpdatedAt = timestamppb.Now()
	task.Status = models.TaskStatus_Ready
	err := r.DB.UpdateTask(task)
	if err != nil {
		logger.Error("save task", zap.String("id", task.Id), zap.Error(err))
		return err
	}
	return nil
}
