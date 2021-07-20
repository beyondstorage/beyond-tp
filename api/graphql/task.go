package graphql

import (
	"context"

	"github.com/beyondstorage/go-toolbox/zapcontext"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/beyondstorage/beyond-tp/models"
)

// runTask handle publish task and update
func (r *mutationResolver) runTask(ctx context.Context, task *models.Task) (err error) {
	gc := GinContextFrom(ctx)
	logger := zapcontext.FromGin(gc)

	task.UpdatedAt = timestamppb.Now()
	task.Status = models.TaskStatus_Running

	txn := r.DB.NewTxn(true)
	err = r.DB.UpdateTask(txn, task)
	if err != nil {
		logger.Error("save task", zap.String("id", task.Id), zap.Error(err))
		txn.Discard()
		return err
	}

	for _, staffId := range task.StaffIds {
		err = r.DB.InsertStaffTask(txn, staffId, task.Id)
		if err != nil {
			logger.Error("insert staff task",
				zap.String("task", task.Id), zap.String("staff", staffId), zap.Error(err))
			txn.Discard()
			return err
		}
	}
	return txn.Commit()
}
