package main

func getAllBooks() []Book {
	books := []Book{}
	result := db.Find(&books)

	if result.Error != nil {
		panic(result.Error)
	}

	return books
}

func getBookByID(id int64) (Book, error) {
	book := Book{}
	err := db.First(&book, id).Error

	if err != nil {
		return book, err
	}

	return book, nil
}

func createNewBook(book *Book) {
	db.Create(&book)
}

func updateBookByID(book *Book, updatedBook *Book) {
	db.Model(&book).Updates(&updatedBook)
}

func deleteBookByID(book *Book) {
	db.Delete(&book)
}
