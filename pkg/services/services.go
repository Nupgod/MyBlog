package services

import (
	"ex01"
	"ex01/pkg/repository"
)

type Blog interface {
	CreatePost(data ex01.Post) error
	GetPosts(offset int) (ex01.PostList, error)
	CountPosts() (int, error)
}

type Service struct {
	Blog
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Blog: NewCreateService(repo.Blog),
	}
}
