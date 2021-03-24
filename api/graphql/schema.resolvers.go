package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/aos-dev/go-toolbox/zapcontext"
	"go.uber.org/zap"

	"github.com/aos-dev/dm/models"
)

func (r *mutationResolver) CreateTask(ctx context.Context, input *CreateTask) (*models.Task, error) {
	gc := GinContextFrom(ctx)
	logger := zapcontext.FromGin(gc)
	db := r.DB

	t, err := input.FormatTask()
	if err != nil {
		return nil, err
	}

	task := models.NewTask()
	task.Name = input.Name
	if input.Status != nil {
		task.Status = *input.Status
	}

	if err = db.CreateTask(task); err != nil {
		return nil, err
	}

	if err = r.Portal.Publish(ctx, t); err != nil {
		return nil, err
	}

	go func() {
		err := r.Portal.Wait(ctx, t)
		if err != nil {
			logger.Error("task running failed", zap.Error(err))
			return
		}
		logger.Info("task exec succeed", zap.String("task_id", t.Id))
		// set task status into finished async
		err = db.SetTaskStatus(t.Id, models.StatusFinished)
		if err != nil {
			logger.Error("set task status failed", zap.String("task_id", t.Id), zap.Error(err))
		}
	}()
	return task, nil
}

func (r *mutationResolver) DeleteTask(ctx context.Context, input *DeleteTask) (*models.Task, error) {
	db := models.DBFromGin(GinContextFrom(ctx))

	// try to get task first
	task, err := db.GetTask(input.ID)
	if err != nil {
		return nil, err
	}
	// then delete task
	if err = db.DeleteTask(input.ID); err != nil {
		return nil, err
	}
	return task, nil
}

func (r *queryResolver) Task(ctx context.Context, id string) (*models.Task, error) {
	db := models.DBFromGin(GinContextFrom(ctx))
	return db.GetTask(id)
}

func (r *queryResolver) Tasks(ctx context.Context) ([]*models.Task, error) {
	db := models.DBFromGin(GinContextFrom(ctx))
	return db.ListTasks()
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
