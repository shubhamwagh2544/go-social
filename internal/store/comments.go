package store

import (
	"context"
	"database/sql"
)

type Comment struct {
	Id        int64  `json:"id"`
	PostId    int64  `json:"post_id"`
	UserId    int64  `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type CommentStore struct {
	db *sql.DB
}

func (c *CommentStore) GetByPostId(ctx context.Context, postId int64) (*Post, error) {
	return nil, nil
}
