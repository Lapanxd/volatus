package route

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lapanxd/volatus-api/internal/dto"
	"github.com/lapanxd/volatus-api/internal/service"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.RouterGroup, db *gorm.DB) {
	r.GET("/me", func(c *gin.Context) {
		Me(c, db)
	})

	r.GET("/:id", func(c *gin.Context) {
		GetById(c, db)
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

	response := dto.UserOutput{
		ID:       user.ID,
		Username: user.Username,
	}

	c.JSON(http.StatusOK, response)
}

func GetById(c *gin.Context, db *gorm.DB) {
	idParam := c.Param("id")

	idInt, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	id := uint(idInt)

	user, err := service.GetUserById(db, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response := dto.UserOutput{
		ID:       user.ID,
		Username: user.Username,
	}

	c.JSON(http.StatusOK, response)
}
