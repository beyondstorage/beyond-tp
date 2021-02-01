package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func StartServer() (err error) {
	r := gin.New()

	// register middleware here
	r.Use(gin.Recovery()) // recover any panic
	r.Use(gin.Logger())
	r.Use(requestID()) // add request-id

	// register routers here
	r.GET("/ping", ping)

	srv := &http.Server{
		Addr:    ":7487",
		Handler: r,
	}

	go func() {
		// connect service
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// graceful restart
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
	return nil
}

func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
