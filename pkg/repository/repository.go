package repository

import (
	"ex01"
	"github.com/jmoiron/sqlx"
)

type Blog interface {
	CreatePost(data ex01.Post) error
	GetPosts(offset int) (ex01.PostList, error)
	CountPosts() (int, error)
}

type Repository struct {
	Blog
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Blog: NewCreatePostgres(db),
	}
}
