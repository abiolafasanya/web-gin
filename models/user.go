package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"type:varchar(256)" json:"name"`
	Email    string `gorm:"index;type:varchar(256)" json:"email"`
	Password string `gorm:"->;<-;not null" json:"password"`
	Token    string `gorm:"-" json:"token,omitempty"`
}
