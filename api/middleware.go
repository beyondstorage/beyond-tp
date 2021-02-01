package api

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/aos-dev/dm/utils"
)

func requestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := utils.NewUUID()
		c.Set("request_id", requestID)
		c.Writer.Header().Set("x-dm-request-id", requestID) // set uuid in header for frontend use
		log.Println("route:", c.HandlerName(), "@", time.Now(), "request_id:", requestID)
		c.Next()
	}
}
