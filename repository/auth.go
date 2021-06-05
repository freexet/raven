package repository

import (
	"time"

	"github.com/freexet/raven/auth"
)

func (s *Storage) CreateUser(user *auth.User) error {
	result := s.db.Create(user)
	return result.Error
}

func (s *Storage) GetUser(user *auth.User) (*auth.User, error) {
	result := s.db.First(user)
	return user, result.Error
}

func (s *Storage) CreateFailedLoginAttempt(attempt *auth.FailedLoginAttemp) error {
	result := s.db.Create(attempt)
	return result.Error
}

func (s *Storage) GetFailedLoginAttempts(ipAddress string) ([]*auth.FailedLoginAttemp, error) {
	var attempts []*auth.FailedLoginAttemp
	result := s.db.Where("ip_address = ? AND created_at > ?", ipAddress, time.Now().Add((-30)*time.Minute)).Find(&attempts)

	return attempts, result.Error
}
