package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/aos-dev/go-toolbox/zapcontext"

	"github.com/aos-dev/dm/models"
)

func (r *mutationResolver) CreateTask(ctx context.Context, input *CreateTask) (*models.Task, error) {
	gc := GinContextFrom(ctx)
	_ = zapcontext.FromGin(gc)
	db := r.DB

	pt, err := input.FormatTask()
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

	// if not running, return task directly and do nothing
	if !input.Status.IsRunning() {
		return task, nil
	}

	// otherwise, run task
	err = r.runTask(ctx, task, pt)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *mutationResolver) DeleteTask(ctx context.Context, input *DeleteTask) (*models.Task, error) {
	// try to get task first
	task, err := r.DB.GetTask(input.ID)
	if err != nil {
		return nil, err
	}
	// then delete task
	if err = r.DB.DeleteTask(input.ID); err != nil {
		return nil, err
	}
	return task, nil
}

func (r *queryResolver) Task(ctx context.Context, id string) (*models.Task, error) {
	return r.DB.GetTask(id)
}

func (r *queryResolver) Tasks(ctx context.Context) ([]*models.Task, error) {
	return r.DB.ListTasks()
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
