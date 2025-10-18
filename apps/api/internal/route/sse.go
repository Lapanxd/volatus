package route

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lapanxd/volatus-api/internal/sse"
)

func SSERoutes(r *gin.RouterGroup) {
	r.GET("/events", func(c *gin.Context) {
		userID := c.GetUint("user_id")

		c.Header("Content-Type", "text/event-stream")
		c.Header("Cache-Control", "no-cache")
		c.Header("Connection", "keep-alive")

		ch := make(chan sse.SSEEvent, 10)
		sse.Register(userID, ch)
		defer sse.UnRegister(userID)

		log.Printf("[SSE] Client %d connected", userID)

		c.SSEvent("init", "connected")
		c.Writer.Flush()
		
		clientGone := c.Request.Context().Done()

		for {
			select {
			case event, ok := <-ch:
				if !ok {
					return
				}
				c.SSEvent(event.EventType, event.Payload)
				c.Writer.Flush()
			case <-clientGone:

				log.Printf("[SSE] Client %d disconnected", userID)
				return
			}
		}
	})
}
