package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lapanxd/volatus-api/internal/dto"
	"github.com/lapanxd/volatus-api/internal/service"
	"gorm.io/gorm"
)

func AuthRoutes(r *gin.RouterGroup, db *gorm.DB) {
	r.POST("/register", func(c *gin.Context) {
		Register(c, db)
	})

	r.POST("/login", func(c *gin.Context) {
		Login(c, db)
	})
}

func Register(c *gin.Context, db *gorm.DB) {
	var input dto.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	user, err := service.RegisterUser(db, input.Username, input.Password)
	if err != nil {
		c.Error(err)
		return
	}

	response := dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
	}

	c.JSON(http.StatusCreated, response)
}

func Login(c *gin.Context, db *gorm.DB) {
	var input dto.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	user, err := service.AuthenticateUser(db, input.Username, input.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	tokenString, err := service.GenerateJWT(user.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	c.JSON(http.StatusOK, dto.LoginResponse{
		AccessToken: tokenString,
	})
}
