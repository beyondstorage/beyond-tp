package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/aos-dev/dm/models"
)

func (r *mutationResolver) CreateTask(ctx context.Context, input *CreateTask) (*models.Task, error) {
	db := models.DBFromGin(GinContextFrom(ctx))

	task := models.NewTask()
	task.Name = input.Name
	if input.Status != nil {
		task.Status = *input.Status
	}
	err := db.CreateTask(task)
	if err != nil {
		return nil, err
	}
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
