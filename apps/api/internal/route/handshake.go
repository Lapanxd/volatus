package route

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
}

func HandshakeInit(c *gin.Context, db *gorm.DB) {
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

	if !CheckIfUserExists(db, req.ToUserID) {
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
	}
	storeMutex.Unlock()

	c.JSON(http.StatusOK, dto.InitResponse{
		SessionID: sessionID,
	})
}

func CheckIfUserExists(db *gorm.DB, userID uint) bool {
	_, err := service.GetUserById(db, userID)
	return err == nil
}
