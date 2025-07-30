package models

import (
	"time"
)

type ContactMessage struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
	Read    bool   `json:"read" gorm:"default:false"` 
	CreatedAt time.Time `json:"created_at"`
}

func (ContactMessage) TableName() string {
	return "contact_messages"
}
