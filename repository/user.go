package repository

import "github.com/freexet/raven/auth"

func (s *Storage) CreateUser(user *auth.User) error {
	result := s.db.Create(user)
	return result.Error
}

func (s *Storage) GetUser(user *auth.User) (*auth.User, error) {
	result := s.db.Where(user).First(user)
	return user, result.Error
}
