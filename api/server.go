package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/aos-dev/noah/task"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/aos-dev/dm/api/graphql"
	"github.com/aos-dev/dm/models"
)

// Server handle configs to start a server
type Server struct {
	Host string
	Port int

	Debug bool

	Logger *zap.Logger
	DB     *models.DB
	Portal *task.Portal
}

// Start a HTTP server
func (s *Server) Start() error {
	// set gin mode instead of env GIN_MODE
	mode := gin.ReleaseMode
	if s.Debug {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	r := gin.New()
	logger := s.Logger

	// register middleware here
	r.Use(setRequestID(), setLoggerWithReqID(logger)) // set requestID and logger
	r.Use(logRequestInfo())                           // log request info
	r.Use(ginRecovery())                              // recover any panic
	r.Use(cors.Default())                             // cors all allowed

	r.Static("web", "ui/build/web")
	r.LoadHTMLGlob("ui/build/web/index.html")

	// register routers here
	r.GET("/ping", ping)
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// register routers for graphql
	gqlServer := graphql.Server{
		Path:   "/graphql",
		Debug:  s.Debug,
		DB:     s.DB,
		Portal: s.Portal,
	}
	gqlServer.RegisterRouter(r)

	endpoint := fmt.Sprintf("%s:%d", s.Host, s.Port)
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
