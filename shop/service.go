package shop

import (
	"errors"

	"github.com/google/uuid"
)

type Service interface {
	CreateShop(userId, name, desc, country, region, city string) (*Shop, error)
}

type Repository interface {
	CreateShop(shop *Shop) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) CreateShop(userId, name, desc, country, region, city string) (*Shop, error) {
	shop := &Shop{
		ID:          uuid.NewString(),
		UserID:      userId,
		Name:        name,
		Description: desc,
		Country:     country,
		Region:      region,
		City:        city,
	}
	if err := s.r.CreateShop(shop); err != nil {
		return nil, errors.New("error creating shop: failed creating shop")
	}

	return shop, nil
}
