package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/beyondstorage/beyond-tp/task"
	"github.com/beyondstorage/go-toolbox/zapcontext"
)

func (r *mutationResolver) CreateTask(ctx context.Context, input *CreateTask) (*Task, error) {
	gc := GinContextFrom(ctx)
	_ = zapcontext.FromGin(gc)

	t := task.NewTask(input.Name, parseTaskType(input.Type))
	t.Storages = input.Storages

	err := r.ts.InsertTask(ctx, t)
	if err != nil {
		return nil, fmt.Errorf("insert task: %w", err)
	}
	return formatTask(t), nil
}

func (r *mutationResolver) DeleteTask(ctx context.Context, input *DeleteTask) (*Task, error) {
	t, err := r.ts.GetTask(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("get task: %w", err)
	}
	if t == nil {
		return nil, fmt.Errorf("task not exist")
	}

	err = r.ts.DeleteTask(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	return formatTask(t), nil
}

func (r *mutationResolver) RunTask(ctx context.Context, id string) (*Task, error) {
	panic("implement me!")
}

func (r *mutationResolver) CreateService(ctx context.Context, input *CreateService) (*Service, error) {
	panic("implement me!")
}

func (r *mutationResolver) DeleteService(ctx context.Context, input *DeleteService) (*Service, error) {
	panic("implement me!")
}

func (r *queryResolver) Task(ctx context.Context, id string) (*Task, error) {
	t, err := r.ts.GetTask(ctx, id)
	if err != nil {
		return nil, err
	}
	return formatTask(t), nil
}

func (r *queryResolver) Tasks(ctx context.Context) ([]*Task, error) {
	panic("implement me!")
}

func (r *queryResolver) Service(ctx context.Context, typeArg ServiceType, name string) (*Service, error) {
	panic("implement me!")
}

func (r *queryResolver) Services(ctx context.Context, typeArg *ServiceType) ([]*Service, error) {
	panic("implement me!")
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
