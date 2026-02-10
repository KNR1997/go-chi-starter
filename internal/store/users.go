package store

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"uniqueIndex;not null" json:"username"`
	Email     string         `gorm:"uniqueIndex;not null" json:"email"`
	Password  string         `json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type UsersStore struct {
	db *gorm.DB
}

func (s UsersStore) Create(ctx context.Context, user *User) error {
	return s.db.WithContext(ctx).Create(user).Error
}
