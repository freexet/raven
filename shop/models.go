package shop

import (
	"github.com/freexet/raven/auth"
	"gorm.io/gorm"
)

type Shop struct {
	gorm.Model
	ID          string `gorm:"type:uuid;primarykey"`
	UserID      string `gorm:"type:uuid"`
	Name        string `gorm:"type:varchar(32);not null"`
	Description string
	Country     string `gorm:"type:varchar(64)"`
	Region      string `gorm:"type:varchar(64)"`
	City        string `gorm:"type:varchar(64)"`
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Shop{})

	db.Migrator().CreateConstraint(&auth.User{}, "Shops")
	db.Migrator().CreateConstraint(&auth.User{}, "fk_users_shops")
}
