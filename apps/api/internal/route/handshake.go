package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lapanxd/volatus-api/internal/dto"
	"github.com/lapanxd/volatus-api/internal/service"
	"gorm.io/gorm"
)

func HandshakeRoutes(r *gin.RouterGroup, db *gorm.DB) {
	r.POST("/init", func(c *gin.Context) {
		HandshakeInit(c, db)
	})

	r.POST("/response", HandshakeResponse)

	r.GET("/pending", GetPendingHandshake)
}

func HandshakeInit(c *gin.Context, db *gorm.DB) {
	var req dto.InitInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	fromUserID := c.GetUint("user_id")
	if !service.CheckIfUserExists(db, req.ToUserID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "target user does not exist"})
		return
	}

	sessionID, err := service.InitHandshake(fromUserID, req.ToUserID, req.SDPOffer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.InitOutput{SessionID: sessionID})
}

func HandshakeResponse(c *gin.Context) {
	var req dto.ResponseInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	userID := c.GetUint("user_id")
	session, err := service.RespondHandshake(userID, req.SessionID, req.Accepted)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if session == nil {
		// return 200 even if handshake is not accepted
		c.Status(http.StatusOK)
		return
	}

	c.Status(http.StatusOK)
}

func GetPendingHandshake(c *gin.Context) {
	userID := c.GetUint("user_id")
	pending := service.GetPendingHandshakes(userID)
	c.JSON(http.StatusOK, dto.PendingOutput{PendingSessions: pending})
}
