package ports

import "go-forum/internal/core/domain"

type PostService interface {
	GetPosts() ([]domain.Post, error)
	GetPostByID(id int) (domain.Post, error)
	GetPostByName(name string) (domain.Post, error)
	CreatePost(name string, text string) error
	DeletePost(id int) error
}

type PostRepository interface {
	GetPosts() ([]domain.Post, error)
	GetPostByID(id int) (domain.Post, error)
	GetPostByName(name string) (domain.Post, error)
	CreatePost(name string, text string) error
	DeletePost(id int) error
}

type CommentService interface {
	CommentPost()
	CommentReply()
}

type CommentRepository interface {
	CommentPost()
	CommentReply()
}

type UserService interface {
	GetUserByID(id int) (domain.User, error)
	CreateUser(name string, email string, password string) (domain.User, error)
	DeleteUser(id int) error
	UpdateUser(id int, name string, email string, password string) (domain.User, error)
	GetAllUsers() ([]domain.User, error)
}

type UserRepository interface {
	GetUserByID(id int) (domain.User, error)
	CreateUser(name string, email string, password string) (domain.User, error)
	DeleteUser(id int) error
	UpdateUser(id int, name string, email string, password string) (domain.User, error)
	GetAllUsers() ([]domain.User, error)
}
