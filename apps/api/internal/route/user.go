package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lapanxd/volatus-api/internal/dto"
	"github.com/lapanxd/volatus-api/internal/service"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.RouterGroup, db *gorm.DB) {
	r.GET("/me", func(c *gin.Context) {
		Me(c, db)
	})
}

func Me(c *gin.Context, db *gorm.DB) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found in context"})
		return
	}
	userID := userIDVal.(uint)

	user, err := service.GetUserById(db, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response := dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
	}

	c.JSON(http.StatusOK, response)
}
