package domain

// Postgres
type Post struct {
	Title       string
	Subject     string
	Text        string
	PicturePath string
}

type User struct {
	Name        string
	PicturePath string
}

type Comment struct {
	Text string
}
