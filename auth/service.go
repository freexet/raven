package auth

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(username, password string) (*User, error)
	Login(username, password string) (*User, error)
}

type Repository interface {
	CreateUser(user *User) error
	GetUser(user *User) (*User, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Register(username, password string) (*User, error) {
	_, err := s.r.GetUser(&User{Username: username})
	if err == nil {
		return nil, errors.New("error registering user: username already exists")
	}

	pwdHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return nil, errors.New("error registering user: failed generating password")
	}

	user := &User{ID: uuid.NewString(), Username: username, PasswordHash: string(pwdHash)}
	err = s.r.CreateUser(user)

	return user, err
}

func (s *service) Login(username, password string) (*User, error) {
	u, err := s.r.GetUser(&User{Username: username})
	if err != nil {
		return nil, errors.New("error login: username not registered")
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("error login: wrong password")
	}

	return u, nil
}
