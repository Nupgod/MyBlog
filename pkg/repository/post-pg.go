package repository

import (
	"ex01"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type CreatePostgres struct {
	db *sqlx.DB
}

func NewCreatePostgres(db *sqlx.DB) *CreatePostgres {
	return &CreatePostgres{db: db}
}

func (r *CreatePostgres) CreatePost(data ex01.Post) error {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, p_text, p_time) VALUES ($1, $2, $3) RETURNING id", postTable)
	row := r.db.QueryRow(query, data.Title, data.Text, data.Time)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (r *CreatePostgres) GetPosts(offset int) (ex01.PostList, error) {
	var posts ex01.PostList
	query := fmt.Sprintf("SELECT title, p_text, p_time FROM %s OFFSET %d LIMIT 3", postTable, offset)
	if err := r.db.Select(&posts, query); err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *CreatePostgres) CountPosts() (int, error) {
	var count int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", postTable)
	row := r.db.QueryRow(query)
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}