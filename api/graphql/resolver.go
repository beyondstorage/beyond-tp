//go:generate go run github.com/99designs/gqlgen ./api/graphql/gqlgen.yml

package graphql

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

var GinCtxKey struct{}

// ginContextFromContext inspired from `https://gqlgen.com/recipes/gin/#accessing-gincontext`
func ginContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(GinCtxKey)
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}
