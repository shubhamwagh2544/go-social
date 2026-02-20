package store

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/lib/pq"
)

type Post struct {
	Id        int64    `json:"id"`
	Content   string   `json:"content"`
	Title     string   `json:"title"`
	Tags      []string `json:"tags"`
	UserId    int64    `json:"user_id"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type PostStore struct {
	db *sql.DB
}

var (
	ErrNotFound = errors.New("resource not found")
)

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	query := `
			insert into posts (content, title, tags, user_id)
			values ($1, $2, $3, $4) 
			returning id, created_at, updated_at
			`

	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Content,
		post.Title,
		pq.Array(post.Tags),
		post.UserId,
	).Scan(
		&post.Id,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostStore) GetById(ctx context.Context, id int64) (*Post, error) {
	query := `
			select id, content, title, tags, user_id, created_at, updated_at
			from posts 
			where id = $1
			`

	// var post Post
	post := &Post{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&post.Id,
		&post.Content,
		&post.Title,
		pq.Array(&post.Tags),
		&post.UserId,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	log.Printf("Post in Store: %+v\n", post)

	// return &post, nil
	return post, nil
}
