package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/aos-dev/go-toolbox/zapcontext"

	"github.com/aos-dev/dm/models"
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

func (r *mutationResolver) CreateIdentity(ctx context.Context, input *CreateIdentity) (*Identity, error) {
	identity := &models.Identity{
		Name:       input.Name,
		Type:       parseIdentityType(input.Type),
		Credential: parseCredentialInput(input.Credential),
		Endpoint:   parseEndpointInput(input.Endpoint),
	}

	// handle transaction manually
	txn := r.DB.NewTxn(true)
	_, err := r.DB.GetIdentity(txn, parseIdentityType(input.Type), input.Name)
	if err == nil {
		txn.Discard()
		return nil, fmt.Errorf("record with type: %s, name: %s, %w", input.Type.String(), input.Name, models.ErrAlreadyExists)
	}

	if errors.Is(err, models.ErrNotFound) {
		if err = r.DB.InsertIdentity(nil, identity); err != nil {
			txn.Discard()
			return nil, err
		}
		if err = txn.Commit(); err != nil {
			return nil, err
		}
		return formatIdentity(identity), nil
	}

	txn.Discard()
	return nil, err
}

func (r *mutationResolver) DeleteIdentity(ctx context.Context, input *DeleteIdentity) (*Identity, error) {
	// handle transaction manually
	txn := r.DB.NewTxn(true)
	// try to get identity first
	id, err := r.DB.GetIdentity(txn, parseIdentityType(input.Type), input.Name)
	if err != nil {
		txn.Discard()
		return nil, err
	}
	// then delete identity
	if err = r.DB.DeleteIdentity(txn, parseIdentityType(input.Type), input.Name); err != nil {
		txn.Discard()
		return nil, err
	}

	if err = txn.Commit(); err != nil {
		return nil, err
	}
	return formatIdentity(id), nil
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

func (r *queryResolver) Identities(ctx context.Context, typeArg *IdentityType) ([]*Identity, error) {
	var idType *models.IdentityType = nil
	if typeArg != nil {
		t := parseIdentityType(*typeArg)
		idType = &t
	}

	ids, err := r.DB.ListIdentity(idType)
	if err != nil {
		return nil, err
	}
	return formatIdentities(ids), nil
}

func (r *queryResolver) Identity(ctx context.Context, typeArg IdentityType, name string) (*Identity, error) {
	id, err := r.DB.GetIdentity(nil, parseIdentityType(typeArg), name)
	if err != nil {
		return nil, err
	}
	return formatIdentity(id), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
