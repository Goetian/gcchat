package db

type Repository interface {
	Migrate() error
	Create(msg Message) (*Message, error)
	All() ([]Message, error)
	GetByAuthor(author string) (*Message, error)
	Update(id int64, updated Message) (*Message, error)
	Delete(id int64) error
}
