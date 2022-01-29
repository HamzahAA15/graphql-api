package book

import (
	"database/sql"
	"log"
	"sirclo/gql/entities"
)

type Repository_Book struct {
	db *sql.DB
}

func NewRepositoryBook(db *sql.DB) *Repository_Book {
	return &Repository_Book{db}
}

//get Books
func (r *Repository_Book) GetBooks() ([]entities.Book, error) {
	var Books []entities.Book
	results, err := r.db.Query("select id, name, author, publisher, year from books")
	if err != nil {
		log.Fatalf("Error")
	}

	defer results.Close()

	for results.Next() {
		var Book entities.Book

		err = results.Scan(&Book.Id, &Book.Name, &Book.Author, &Book.Publisher, &Book.Year)
		if err != nil {
			log.Fatalf("Error")
		}

		Books = append(Books, Book)
	}
	return Books, nil
}

//get book
func (r *Repository_Book) GetBook(id int) (entities.Book, error) {
	var book entities.Book
	results := r.db.QueryRow("select id, name, author, publisher, year from books WHERE id = ? order by id asc", id)
	// if err != nil {
	// 	log.Fatalf("Error")
	// }

	err := results.Scan(&book.Id, &book.Name, &book.Author, &book.Publisher, &book.Year)
	if err != nil {
		return book, err
	}
	return book, nil
}

//create book
func (r *Repository_Book) CreateBook(book entities.Book) (entities.Book, error) {
	query := `INSERT INTO Books (name, author, publisher, year) VALUES (?, ?, ?, ?)`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return book, err
	}

	_, err = statement.Exec(book.Name, book.Author, book.Publisher, book.Year)
	if err != nil {
		return book, err
	}

	return book, nil
}

//update book
func (r *Repository_Book) UpdateBook(book entities.Book) (entities.Book, error) {
	query := `UPDATE books SET name = ?, author = ?, publisher = ?, year = ? WHERE id = ?`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return book, err
	}

	defer statement.Close()

	_, err = statement.Exec(book.Name, book.Author, book.Publisher, book.Year, book.Id)
	if err != nil {
		return book, err
	}

	return book, nil
}

//delete book
func (r *Repository_Book) DeleteBook(book entities.Book) (entities.Book, error) {
	query := `DELETE FROM books WHERE id = ?`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return book, err
	}

	defer statement.Close()

	_, err = statement.Exec(book.Id)
	if err != nil {
		return book, err
	}

	return book, nil
}
