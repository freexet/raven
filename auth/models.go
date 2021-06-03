package auth

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           string `gorm:"type:varchar(64);primarykey"`
	Username     string `gorm:"type:varchar(32);unique_index"`
	PasswordHash string `gorm:"not null"`
	Token        string `gorm:"-"`
	SecretKey    string
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
