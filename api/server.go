package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ServerConfig handle configs to start a server
type ServerConfig struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`

	Debug bool `json:"debug" yaml:"debug"`

	Logger *zap.Logger
}

// StartServer start a HTTP server
func StartServer(cfg ServerConfig) (err error) {
	// set gin mode instead of env GIN_MODE
	mode := gin.ReleaseMode
	if cfg.Debug {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	r := gin.New()
	logger := cfg.Logger

	// register middleware here
	r.Use(setRequestID(), setLoggerWithReqID(logger)) // set requestID and logger
	r.Use(logRequestInfo())                           // log request info
	r.Use(ginRecovery(cfg.Debug))                     // recover any panic

	// register routers here
	r.GET("/ping", ping)

	endpoint := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	srv := &http.Server{
		Addr:    endpoint,
		Handler: r,
	}

	go func() {
		// connect service
		logger.Info("server now started", zap.String("endpoint", endpoint))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("listen: %s", zap.Error(err))
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
