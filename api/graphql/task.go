package graphql

import (
	"context"

	"github.com/aos-dev/go-toolbox/zapcontext"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/aos-dev/dm/models"
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
