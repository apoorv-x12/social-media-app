package store

import (
	"context"
	"database/sql"
)

type User struct {
}
type UserRepository interface {
	Create(context.Context, *User) error
}
type PgSqlUserRepository struct {
	db *sql.DB
}

func (repo *PgSqlUserRepository) Create(ctx context.Context, user *User) error {

	return nil	

}
