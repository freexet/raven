package repository

import "github.com/freexet/raven/auth"

func (s *Storage) CreateUser(user *auth.User) error {
	result := s.db.Create(user)
	return result.Error
}
