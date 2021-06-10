package repository

import "github.com/freexet/raven/shop"

func (s *Storage) CreateShop(shop *shop.Shop) error {
	result := s.db.Create(shop)
	return result.Error
}
