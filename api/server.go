package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/aos-dev/go-toolbox/zapcontext"
	"github.com/aos-dev/noah/proto"
	"github.com/aos-dev/noah/task"
	"github.com/gin-gonic/gin"
	protobuf "github.com/golang/protobuf/proto"
	"github.com/google/uuid"
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

	r.Static("assets", "ui/dist")

	// register routers here
	r.GET("/ping", ping)

	// register routers for graphql
	graphql.RegisterRouter(r, "/graphql", s.DB, s.Debug)

	// register routers for tasks
	taskGroup := r.Group("/task")
	{
		taskGroup.POST("/copy", s.copyTask)
	}

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

func (s *Server) copyTask(c *gin.Context) {
	logger := zapcontext.FromGin(c)

	copyFileJob := &proto.CopyDir{
		Src:       0,
		Dst:       1,
		SrcPath:   "",
		DstPath:   "",
		Recursive: true,
	}
	content, err := protobuf.Marshal(copyFileJob)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	copyFileTask := &proto.Task{
		Id: uuid.NewString(),
		Endpoints: []*proto.Endpoint{
			{Type: "fs", Pairs: []*proto.Pair{{Key: "work_dir", Value: "/tmp/b/"}}},
			{Type: "fs", Pairs: []*proto.Pair{{Key: "work_dir", Value: "/tmp/c/"}}},
		},
		Job: &proto.Job{
			Id:      uuid.NewString(),
			Type:    task.TypeCopyDir,
			Content: content,
		},
	}

	logger.Info("before first publish")
	err = s.Portal.Publish(c, copyFileTask)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusOK, "copy published")
}
