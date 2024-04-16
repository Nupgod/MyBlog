package services

import (
	"ex01"
	"ex01/pkg/repository"
)

type CreateService struct {
	repo repository.Blog
}

func NewCreateService(repo repository.Blog) *CreateService {
	return &CreateService{repo: repo}
}

func (s *CreateService) CreatePost(data ex01.Post) error {
	return s.repo.CreatePost(data)
}

func (s *CreateService) GetPosts(offset int) (ex01.PostList, error) {
	return s.repo.GetPosts(offset)
}
func (s *CreateService) CountPosts() (int, error) {
	return s.repo.CountPosts()
}