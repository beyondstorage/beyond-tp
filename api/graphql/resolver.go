//go:generate go run github.com/99designs/gqlgen ./api/graphql/gqlgen.yml

package graphql

import (
	"context"

	"github.com/aos-dev/dm/models"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

// mustDBHandlerFrom inspired from `https://gqlgen.com/recipes/gin/#accessing-gincontext`
func mustDBHandlerFrom(ctx context.Context) *models.DBHandler {
	v := ctx.Value(models.DBCtxKey)
	if v == nil {
		panic("could not retrieve DBHandler")
	}

	return v.(*models.DBHandler)
}
