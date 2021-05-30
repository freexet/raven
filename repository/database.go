package repository

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func New() (*Storage, error) {
	s := new(Storage)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	dBase, _ := db.DB()
	dBase.SetConnMaxIdleTime(10)

	s.db = db

	return s, nil
}

func (s *Storage) GetDB() *gorm.DB {
	return s.db
}

func (s *Storage) Begin() *gorm.DB {
	return s.db.Begin()
}
