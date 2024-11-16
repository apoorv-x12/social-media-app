package store

import "database/sql"

type Storage struct {
	Users UserRepository
	Posts PostRepository
}

func NewStorage(db *sql.DB) (*Storage,error) {
 
	return &Storage{
		Users: &PgSqlUserRepository{db},
		Posts: &PgSqlPostRepository{db},       
	},
	nil
}