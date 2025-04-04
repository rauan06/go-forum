package ports

import "go-forum/internal/core/domain"

type PostService interface {
	GetPosts() ([]domain.Post, error)
	GetPostByID(id int) (domain.Post, error)
	CreatePost(id int, name string, text string) error
	DeletePost(id int) error
}

type PostRepository interface {
	GetPosts() ([]domain.Post, error)
	GetPostByID(id int) (domain.Post, error)
	CreatePost(id int, name string, text string) error
	DeletePost(id int) error
}

type CommentService interface {
	CommentPost()
	CommentReply()
}
type CommentRepository interface{}
