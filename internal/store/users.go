package store

import (
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrDuplicateEmail    = errors.New("a user with that email already exists")
	ErrDuplicateUsername = errors.New("a user with that username already exists")
)

type User struct {
	ID           int64          `json:"id"`
	Username     string         `gorm:"uniqueIndex;not null" json:"username"`
	Email        string         `gorm:"uniqueIndex;not null" json:"email"`
	Password     password       `gorm:"-" json:"-"` // not stored directly
	PasswordHash []byte         `gorm:"not null" json:"-"`
	CreatedAt    time.Time      `json:"created_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	IsActive     bool           `json:"is_active"`
	RoleID       int64          `json:"role_id"`
	Role         Role           `json:"role"`
}

type password struct {
	text *string
	hash []byte
}

func (u *User) SetPassword(plain string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.PasswordHash = hash
	return nil
}

func (u *User) CheckPassword(plain string) error {
	return bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(plain))
}

type UsersStore struct {
	db *gorm.DB
}

func (s UsersStore) Create(ctx context.Context, user *User) error {
	return s.db.WithContext(ctx).Create(user).Error
}

func (s *UsersStore) GetByID(ctx context.Context, id int64) (*User, error) {
	var user User

	err := s.db.WithContext(ctx).
		First(&user, id).
		Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UsersStore) GetByEmail(ctx context.Context, email string) (*User, error) {
	var user User

	err := s.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}

	return &user, err
}
