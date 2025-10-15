package model

import (
	"gorm.io/gorm"
)

type PendingNotification struct {
	gorm.Model
	UserID    uint
	EventType string
	Payload   string
}
