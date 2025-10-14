package route

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lapanxd/volatus-api/internal/dto"
)

type HandshakeSession struct {
	FromUserID uint
	ToUserID   uint
}

// persist sessions in database when needed (redis or postgres)
var (
	handshakeStore = make(map[string]HandshakeSession)
	storeMutex     = sync.Mutex{}
)

func HandshakeRoutes(r *gin.RouterGroup) {
	r.POST("/init", HandshakeInit)
}

func HandshakeInit(c *gin.Context) {
	var req dto.InitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	fromUserIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing user"})
		return
	}
	fromUserID := fromUserIDVal.(uint)

	sessionID := uuid.NewString()

	storeMutex.Lock()
	handshakeStore[sessionID] = HandshakeSession{
		FromUserID: fromUserID,
		ToUserID:   req.ToUserID,
	}
	storeMutex.Unlock()

	c.JSON(http.StatusOK, dto.InitResponse{
		SessionID: sessionID,
	})
}
