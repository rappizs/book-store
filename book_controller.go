package main

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
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

	result, err := getBookByID(id)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	json.NewEncoder(w).Encode(result)
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

	err := createNewBook(book)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 0, 36)

	book, err := getBookByID(id)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	json.NewDecoder(r.Body).Decode(&book)
	err = updateBookByID(book)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 0, 36)

	book, err := getBookByID(id)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	err = deleteBookByID(book)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
}
