package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/beyondstorage/beyond-tp/api/graphql"
	"github.com/beyondstorage/beyond-tp/api/ui"
	"github.com/beyondstorage/beyond-tp/task"
)

// Server handle configs to start a server
type Server struct {
	Host string
	Port int

	DevMode bool

	Logger *zap.Logger
	ts     *task.Server
}

// Start a HTTP server
func (s *Server) Start() error {
	// set gin mode instead of env GIN_MODE
	mode := gin.ReleaseMode
	if s.DevMode {
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

	// register routers here
	r.GET("/ping", ping)
	r.StaticFS("/ui", http.FS(ui.UI))

	// redirect / to /ui, since /ui is the real homepage of website
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/ui")
	})

	r.NoRoute(func(c *gin.Context) {
		c.Status(http.StatusNotFound)
	})

	// register routers for graphql
	gqlServer := graphql.Server{
		Path:  "/graphql",
		Debug: s.DevMode,
		Task:  s.ts,
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
