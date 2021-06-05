package auth

import (
	"time"

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

type FailedLoginAttemp struct {
	ID        string `gorm:"type:varchar(64);primarykey"`
	IPAddress string `gorm:"type:varchar(32);not null;column:ip_address"`
	CreatedAt time.Time
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&FailedLoginAttemp{})
}
