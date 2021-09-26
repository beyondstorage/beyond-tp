package graphql

import (
	"context"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	"github.com/beyondstorage/beyond-tp/models"
)

const ginCtxKey = "gin_in_ctx"

type Server struct {
	Path  string
	Debug bool

	DB *models.DB
}

func (s *Server) RegisterRouter(r *gin.Engine) {
	// register routers for graphql
	gqlGroup := r.Group(s.Path)
	// register db into gin context, then set gin ctx into context
	gqlGroup.Use(WithInGinCtx())
	// enable playground only in debug mode
	if s.Debug {
		playGroundHandler := playground.Handler("GraphQL playground", s.Path)
		gqlGroup.GET("", gin.WrapF(playGroundHandler))
	}
	gplHandler := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: &Resolver{
		DB:      s.DB,
	}}))
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
