package store

import (
	"gorm.io/gorm"
)

type Storage struct {
	Posts PostsStore
	Users UsersStore
}

func NewStorage(db *gorm.DB) Storage {
	return Storage{
		Posts: PostsStore{db},
		Users: UsersStore{db},
	}
}
