package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/aos-dev/go-toolbox/zapcontext"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/aos-dev/dm/api/graphql"
)

// Config handle configs to start a server
type Config struct {
	Host string
	Port int

	Debug  bool
	DBPath string
}

// StartServer start a HTTP server
func StartServer(ctx context.Context, cfg Config) (err error) {
	logger := zapcontext.From(ctx)

	// set gin mode instead of env GIN_MODE
	mode := gin.ReleaseMode
	if cfg.Debug {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	r := gin.New()

	r.Use(WithRequestID())
	r.Use(WithLogger(logger))
	r.Use(WithRecovery()) // recover any panic
	r.Static("assets", "ui/dist")

	// register routers here
	r.GET("/ping", ping)

	// register routers for graphql
	graphql.RegisterRouter(ctx, r, "/graphql", cfg.Debug)

	endpoint := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	srv := &http.Server{
		Addr:    endpoint,
		Handler: r,
	}

	go func() {
		// connect service
		logger.Info("server now started", zap.String("endpoint", endpoint))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("listen:", zap.Error(err))
		}
	}()

	// graceful restart
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Warn("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown:", zap.Error(err))
	}
	logger.Warn("Server exiting")
	return nil
}

func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
