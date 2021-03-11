package graphql

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	"github.com/aos-dev/dm/models"
)

func RegisterRouter(ctx context.Context, r *gin.Engine, relPath string, debug bool) {
	// register routers for graphql
	group := r.Group(relPath)
	// enable playground only in debug mode
	if debug {
		playGroundHandler := playground.Handler("GraphQL playground", relPath)
		group.GET("", gin.WrapF(playGroundHandler))
	}

	db := models.DBFromContext(ctx)

	srv := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: &Resolver{}}))

	// Set db into graphql.
	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		ctx = models.DbIntoContext(ctx, db)

		return next(ctx)
	})
	group.POST("", gin.WrapH(srv))
}
