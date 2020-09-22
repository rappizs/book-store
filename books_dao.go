package main

import (
	"database/sql"
)

func getAllBooks() []Book {
	rows, err := db.Query("SELECT id, title, isbn FROM books")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := []Book{}

	for rows.Next() {
		book := Book{}
		err = rows.Scan(&book.ID, &book.Title, &book.Isbn)
		if err != nil {
			panic(err)
		}
		result = append(result, book)
	}

	return result
}

func getBookByID(id int64) (Book, error) {
	row := db.QueryRow("SELECT id, title, isbn FROM books WHERE id = $1", id)
	book := Book{}

	switch err := row.Scan(&book.ID, &book.Title, &book.Isbn); err {
	case sql.ErrNoRows:
		return book, err
	case nil:
		return book, nil
	default:
		panic(err)
	}
}

func createNewBook(book Book) error {
	statement := `
		INSERT INTO books (title, isbn)
		Values($1, $2)`

	_, err := db.Exec(statement, book.Title, book.Isbn)
	if err != nil {
		return err
	}
	return nil
}

func updateBookByID(book Book) error {
	statement := `
		UPDATE books 
		SET title = $2, isbn = $3
		WHERE id = $1`

	_, err := db.Exec(statement, book.ID, book.Title, book.Isbn)
	if err != nil {
		return err
	}
	return nil
}

func deleteBookByID(book Book) error {
	statement := `
		DELETE FROM books 
		WHERE id = $1`

	_, err := db.Exec(statement, book.ID)
	if err != nil {
		return err
	}
	return nil
}
