package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully connected to %v\n", dbname)

	startRouter()
}

func startRouter(){
	r := mux.NewRouter()

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PATCH")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	fmt.Printf("Application listening on localhost, port: %v\n", serverPort)
	log.Fatal(http.ListenAndServe(":" + serverPort, r))
}
