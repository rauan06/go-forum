package services

import (
	"go-forum/internal/core/domain"
	"go-forum/internal/core/ports"
)

type PostService struct {
	repo ports.PostRepository
}

func (s *PostService) NewPostService(repo ports.PostRepository) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (s *PostService) GetPosts() ([]domain.Post, error) {
	return s.repo.GetPosts()
}

func (s *PostService) GetPostByID(id int) (domain.Post, error) {
	return s.repo.GetPostByID(id)
}

func (s *PostService) CreatePost(id int, name string, text string) error {
	return s.repo.CreatePost(id, name, text)
}

func (s *PostService) DeletePost(id int) error {
	return s.repo.DeletePost(id)
}
