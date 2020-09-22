package main

//Book is a representation of a real book
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author is a representation of an author
type Author struct {
	FirstName string `json:"firstname"`
	LatName   string `json:"lastname"`
}
