package service

import (
	"errors"

	"github.com/lapanxd/volatus-api/internal/model"
	"gorm.io/gorm"
)

func GetUserById(db *gorm.DB, userId uint) (*model.User, error) {
	var user model.User
	if err := db.First(&user, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func CheckIfUserExists(db *gorm.DB, userID uint) bool {
	_, err := GetUserById(db, userID)
	return err == nil
}
