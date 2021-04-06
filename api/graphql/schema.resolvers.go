package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/aos-dev/dm/models"
	"github.com/aos-dev/go-toolbox/zapcontext"
)

func (r *mutationResolver) CreateTask(ctx context.Context, input *CreateTask) (*Task, error) {
	gc := GinContextFrom(ctx)
	_ = zapcontext.FromGin(gc)
	db := r.DB

	task := models.NewTask(input.Name, parseTaskType(input.Type))
	task.Options = parsePairsInput(input.Options)
	task.Storages = parseStoragesInput(input.Storages)
	// TODO: we need to check the staffs status.
	for _, v := range input.Staffs {
		task.StaffIds = append(task.StaffIds, v.ID)
	}

	if err := db.InsertTask(nil, task); err != nil {
		return nil, err
	}
	return formatTask(task), nil
}

func (r *mutationResolver) DeleteTask(ctx context.Context, input *DeleteTask) (*Task, error) {
	// try to get task first
	task, err := r.DB.GetTask(input.ID)
	if err != nil {
		return nil, err
	}
	// then delete task
	if err = r.DB.DeleteTask(input.ID); err != nil {
		return nil, err
	}
	return formatTask(task), nil
}

func (r *mutationResolver) RunTask(ctx context.Context, id string) (*Task, error) {
	// try to get task first
	task, err := r.DB.GetTask(id)
	if err != nil {
		return nil, err
	}

	if err = r.runTask(ctx, task); err != nil {
		return nil, err
	}
	return formatTask(task), nil
}

func (r *queryResolver) Task(ctx context.Context, id string) (*Task, error) {
	t, err := r.DB.GetTask(id)
	if err != nil {
		return nil, err
	}
	return formatTask(t), nil
}

func (r *queryResolver) Tasks(ctx context.Context) ([]*Task, error) {
	ts, err := r.DB.ListTasks()
	if err != nil {
		return nil, err
	}
	return formatTasks(ts), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
