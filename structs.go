package main

//Book is a representation of a real book
type Book struct {
	//gorm.Model
	ID    uint   `json:"id" gorm:"primaryKey"`
	Isbn  string `json:"isbn" gorm:"not null"`
	Title string `json:"title"  gorm:"not null"`
	//Author *Author `json:"author"`
}

//Author is a representation of an author
type Author struct {
	FirstName string `json:"firstname"`
	LatName   string `json:"lastname"`
}
