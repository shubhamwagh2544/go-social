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
	User      User   `json:"user"`
}

type CommentStore struct {
	db *sql.DB
}

func (c *CommentStore) GetByPostId(ctx context.Context, postId int64) ([]Comment, error) {
	query := `
			select c.id, c.post_id, c.user_id, c.content, c.created_at, u.username, u.id from comments c
			join users u on c.user_id = u.id
			where c.post_id = $1
			order by c.created_at desc
			`
	rows, err := c.db.QueryContext(ctx, query, postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []Comment{}

	for rows.Next() {
		var comment Comment
		comment.User = User{}

		err := rows.Scan(
			&comment.Id,
			&comment.PostId,
			&comment.UserId,
			&comment.Content,
			&comment.CreatedAt,
			&comment.User.Username,
			&comment.User.Id,
		)

		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
