package route

import "C"
import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lapanxd/volatus-api/internal/dto"
	"github.com/lapanxd/volatus-api/internal/service"
	"gorm.io/gorm"
)

type HandshakeSession struct {
	FromUserID uint
	ToUserID   uint
	OfferSDP   string
}

// persist sessions in database when needed (redis or postgres)
var (
	handshakeStore = make(map[string]HandshakeSession)
	storeMutex     = sync.Mutex{}
)

func HandshakeRoutes(r *gin.RouterGroup, db *gorm.DB) {
	r.POST("/init", func(c *gin.Context) {
		HandshakeInit(c, db)
	})

	r.POST("/response", func(c *gin.Context) {
		HandshakeResponse(c, db)
	})

	r.GET("/pending", func(c *gin.Context) {
		GetPendingHandshakeSession(c)
	})
}

func HandshakeInit(c *gin.Context, db *gorm.DB) {
	var req dto.InitInput
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

	if !service.CheckIfUserExists(db, req.ToUserID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "target user does not exist"})
		return
	}

	if req.ToUserID == fromUserID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot handshake with yourself"})
		return
	}

	sessionID := uuid.NewString()

	storeMutex.Lock()
	handshakeStore[sessionID] = HandshakeSession{
		FromUserID: fromUserID,
		ToUserID:   req.ToUserID,
		OfferSDP:   req.SDPOffer,
	}
	storeMutex.Unlock()

	c.JSON(http.StatusOK, dto.InitResponse{
		SessionID: sessionID,
	})
}

func HandshakeResponse(c *gin.Context, db *gorm.DB) {
	var req dto.ResponseInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing user"})
		return
	}
	userID := userIDVal.(uint)

	storeMutex.Lock()
	session, found := handshakeStore[req.SessionID]

	if !found {
		storeMutex.Unlock()
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "session not found"})
		return
	}

	if session.ToUserID != userID {
		storeMutex.Unlock()
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "not your session"})
		return
	}

	if !req.Accepted {
		delete(handshakeStore, req.SessionID)
		storeMutex.Unlock()
		c.JSON(http.StatusOK, "handshake refused")
		return
	}

	delete(handshakeStore, req.SessionID)
	storeMutex.Unlock()

	// todo: notify "from user" that session has been accepted
	c.Status(http.StatusOK)
}

func GetPendingHandshakeSession(c *gin.Context) {

}
