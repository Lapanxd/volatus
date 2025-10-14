package routes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lapanxd/volatus-api/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

func AuthRoutes(r *gin.RouterGroup, db *gorm.DB) {
	r.POST("/register", func(c *gin.Context) {
		Register(c, db)
	})
}

func Register(c *gin.Context, db *gorm.DB) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Error(errors.New("cannot hash password"))
		return
	}

	user := models.User{
		Username: input.Username,
		Password: string(hashedPassword),
	}

	if err := db.Create(&user).Error; err != nil {
		c.Error(errors.New("cannot create user"))
	}

	response := UserResponse{
		ID:       user.ID,
		Username: user.Username,
	}

	c.JSON(http.StatusCreated, response)
}
