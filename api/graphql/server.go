package graphql

import (
	"context"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	"github.com/aos-dev/dm/models"
)

const ginCtxKey = "gin_in_ctx"

func RegisterRouter(r *gin.Engine, relPath string, db *models.DB, debug bool) {
	// register routers for graphql
	gqlGroup := r.Group(relPath)
	// register db into gin context, then set gin ctx into context
	gqlGroup.Use(models.DbIntoGin(db), WithInGinCtx())
	// enable playground only in debug mode
	if debug {
		playGroundHandler := playground.Handler("GraphQL playground", relPath)
		gqlGroup.GET("", gin.WrapF(playGroundHandler))
	}
	gplHandler := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: &Resolver{}}))
	gqlGroup.POST("", gin.WrapH(gplHandler))
}

// WithInGinCtx set gin context into request's context with key ginCtxKey
// inspired from https://gqlgen.com/recipes/gin/#accessing-gincontext
func WithInGinCtx() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), ginCtxKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// GinContextFrom retrieve gin.Context from request's context
// inspired from https://gqlgen.com/recipes/gin/#accessing-gincontext
func GinContextFrom(ctx context.Context) *gin.Context {
	ginContext := ctx.Value(ginCtxKey)
	if ginContext == nil {
		panic("could not retrieve gin.Context")
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		panic("gin.Context has wrong type")
	}
	return gc
}
