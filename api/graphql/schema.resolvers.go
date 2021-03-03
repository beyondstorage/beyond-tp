package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/aos-dev/dm/api/graphql/generated"
	"github.com/aos-dev/dm/models"
	"github.com/google/uuid"
)

func (r *mutationResolver) CreateTask(ctx context.Context, input *models.CreateTask) (*models.Task, error) {
	now := time.Now()
	task := &models.Task{
		ID:        uuid.NewString(), // generate uuid
		Name:      input.Name,
		Status:    models.StatusCreated, // default status: created
		CreatedAt: now,
		UpdatedAt: now,
	}
	if input.Status != nil {
		task.Status = *input.Status
	}
	err := task.Save()
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *mutationResolver) DeleteTask(ctx context.Context, input *models.DeleteTask) (*models.Task, error) {
	task := models.Task{
		ID: input.ID,
	}
	err := task.Delete()
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *queryResolver) Task(ctx context.Context, id string) (*models.Task, error) {
	return models.GetTask(id)
}

func (r *queryResolver) Tasks(ctx context.Context) ([]*models.Task, error) {
	return models.ListTasks()
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
