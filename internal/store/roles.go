package store

import (
	"context"

	"gorm.io/gorm"
)

type Role struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Level       int    `json:"level"`
}

type RoleStore struct {
	db *gorm.DB
}

func (s *RoleStore) GetByName(ctx context.Context, name string) (*Role, error) {
	var role Role

	err := s.db.WithContext(ctx).
		First(&role, name).
		Error
	if err != nil {
		return nil, err
	}

	return &role, nil
}
