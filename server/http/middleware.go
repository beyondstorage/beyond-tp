package http

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/aos-dev/dm/utils"
)

func requestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := utils.NewUUID()
		c.Set("request_id", requestID)
		c.Writer.Header().Set("X-Request-ID", requestID) // set uuid in header for frontend use
		fmt.Println("request_id:", requestID, "@", time.Now())
		c.Next()
	}
}
