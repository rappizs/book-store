package main

import (
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *gorm.DB

func main() {
	initDB()
	startRouter()
}

func startRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PATCH")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	fmt.Printf("Application listening on localhost, port: %v\n", serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, r))
}

func initDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Book{})

	fmt.Printf("Successfully connected to %v\n", dbname)
}
