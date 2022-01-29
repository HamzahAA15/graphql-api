package book

import (
	"sirclo/gql/entities"
)

type RepositoryBook interface {
	GetBooks() ([]entities.Book, error)
	GetBook(id int) (entities.Book, error)
	CreateBook(book entities.Book) (entities.Book, error)
	UpdateBook(Book entities.Book) (entities.Book, error)
	DeleteBook(Book entities.Book) (entities.Book, error)
}
