package store

import "database/sql"

type Storage struct {
	UserRepository UserRepository
	PostRepository PostRepository
}

func NewStorage(db *sql.DB) (*Storage,error) {
 
	return &Storage{
		UserRepository: &PgSqlUserRepository{db},
		PostRepository: &PgSqlPostRepository{db},       
	},
	nil
}