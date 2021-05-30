package auth

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"type:varchar(32);column:username;unique_index"`
	PasswordHash string `gorm:"column:password;not null"`
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
