package store

import (
	"context"
	"database/sql"
)

type Post struct {
	 
}

type PostRepository interface {
	Create(context.Context, *Post) error
}
type PgSqlPostRepository struct {
	db *sql.DB
}

func (repo *PgSqlPostRepository) Create(ctx context.Context, post *Post) error {

	return nil
}
