package main

/**
 * @Author huang
 * @Date 2024-07-31
 * @File: sse.go
 * @Description: 客户端连接后，服务端不断推送数据
 */

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func sseHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			message := fmt.Sprintf("data: The time is %s\n\n", t.Format(time.RFC3339))
			_, err := c.Writer.Write([]byte(message))
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			c.Writer.Flush()
		case <-c.Request.Context().Done():
			return
		}
	}
}

func main() {
	r := gin.Default()
	r.GET("/sse", sseHandler)
	r.Run(":8080")
}
