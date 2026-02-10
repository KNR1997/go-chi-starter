package store

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
	Title   string `json:"title"`
	UserID  int64  `json:"user_id"`
	// Tags      []string       `json:"tags"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type PostsStore struct {
	db *gorm.DB
}

func (s PostsStore) Create(ctx context.Context, post *Post) error {
	return s.db.WithContext(ctx).Create(post).Error
}

// func (s *PostsStore) Create(ctx context.Context, post *Post) error {
// 	query := `
// 		INSERT INTO posts (content, title, user_id, tags)
// 		VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
// 	`

// 	err := s.db.QueryRowContext(
// 		ctx,
// 		query,
// 		post.Content,
// 		post.Title,
// 		post.UserID,
// 		pq.Array(post.Tags),
// 	).Scan(
// 		&post.ID,
// 		&post.CreatedAt,
// 		&post.UpdatedAt,
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
