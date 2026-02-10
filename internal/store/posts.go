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

type PostStore struct {
	db *gorm.DB
}

func (s PostStore) Create(ctx context.Context, post *Post) error {
	return s.db.WithContext(ctx).Create(post).Error
}

func (s *PostStore) GetByID(ctx context.Context, id int64) (*Post, error) {
	var post Post

	err := s.db.WithContext(ctx).
		First(&post, id).
		Error
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (s *PostStore) Delete(ctx context.Context, id int64) error {
	result := s.db.WithContext(ctx).
		Delete(&Post{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func (s *PostStore) Update(ctx context.Context, post *Post) error {
	result := s.db.WithContext(ctx).
		Model(&Post{}).
		Where("id = ?", post.ID).
		Updates(map[string]interface{}{
			"title":   post.Title,
			"content": post.Content,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
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
