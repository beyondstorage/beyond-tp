package graphql

import (
	"context"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	"github.com/aos-dev/dm/models"
)

func RegisterRouter(ctx context.Context, r *gin.Engine, relPath string, db *models.DB, debug bool) {
	// register routers for graphql
	gqlGroup := r.Group(relPath)
	// enable playground only in debug mode
	if debug {
		playGroundHandler := playground.Handler("GraphQL playground", relPath)
		gqlGroup.GET("", gin.WrapF(playGroundHandler))
	}
	gplHandler := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: &Resolver{}}))
	gqlGroup.POST("", gin.WrapH(gplHandler))
}
