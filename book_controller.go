package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := getAllBooks()
	json.NewEncoder(w).Encode(result)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 0, 36)

	book, err := getBookByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(404)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	book := Book{}
	json.NewDecoder(r.Body).Decode(&book)

	fields := reflect.TypeOf(book)
	values := reflect.ValueOf(book)
	for i := 0; i < fields.NumField(); i++ {
		f := fields.Field(i).Name
		v := values.Field(i).Interface()
		if f != "ID" && f != "Author" {
			if v == "" {
				w.WriteHeader(400)
				return
			}
		}
	}

	createNewBook(&book)
	w.WriteHeader(204)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 0, 36)

	book, err := getBookByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(404)
		return
	}

	updatedBook := Book{}
	json.NewDecoder(r.Body).Decode(&updatedBook)
	updateBookByID(&book, &updatedBook)
	w.WriteHeader(204)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 0, 36)

	book, err := getBookByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(404)
		return
	}

	deleteBookByID(&book)
	w.WriteHeader(204)
}
