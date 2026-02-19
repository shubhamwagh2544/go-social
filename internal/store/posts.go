package store

import (
	"context"
	"database/sql"
)

type Post struct {
	Id        int64    `json:"id"`
	UserId    int64    `json:"user_id"`
	Content   string   `json:"content"`
	Title     string   `json:"title"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type PostStore struct {
	db *sql.DB
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	// sql op
	return nil
}
