package http

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/beyondstorage/beyond-tp/http/graphql"
	"github.com/beyondstorage/beyond-tp/http/ui"
	"github.com/beyondstorage/beyond-tp/task"
)

// Server handle configs to start a server
type Server struct {
	addr   string
	logger *zap.Logger
	ts     *task.Server
}

type Config struct {
	Addr   string
	Logger *zap.Logger
	Task   *task.Server
}

func New(c *Config) (s *Server) {
	s = &Server{
		addr:   c.Addr,
		logger: c.Logger,
		ts:     c.Task,
	}
	return
}

// Serve a HTTP server
func (s *Server) Serve() error {
	// FIXME: we will add global dev mode.
	gin.SetMode(gin.DebugMode)

	r := gin.New()
	logger := s.logger

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
		Debug: true,
		Task:  s.ts,
	}
	gqlServer.RegisterRouter(r)

	srv := &http.Server{
		Addr:    s.addr,
		Handler: r,
	}

	go func() {
		// connect service
		logger.Info("server now started", zap.String("addr", s.addr))
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
