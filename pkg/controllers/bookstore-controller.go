package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/theaveas/gCRUD/pkg/storage"
	"github.com/theaveas/gCRUD/pkg/utils"
)

var NewBook storage.Book

func GetAllBooks(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "pkglication/json")
	books := storage.GetAllBooks()
	response, _ := json.Marshal(books)
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}

func GetBookById(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "pkglication/json")
	params := mux.Vars(req)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Printf("error while parsing ~ %s", err)
	}

	book, _ := storage.GetBookById(ID)
	response, _ := json.Marshal(book)
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}

func CreateBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "pkglication/json")
	newBook := &storage.Book{}
	utils.ParseBody(req, newBook)
	b := newBook.CreateBook()
	response, _ := json.Marshal(b)
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}

func DeleteBookById(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "pkglication/json")
	params := mux.Vars(req)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Printf("error while parsing ~ %s", err)
	}

	book := storage.DeleleBookById(ID)
	response, _ := json.Marshal(book)
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}

func UpdateBookById(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "pkglication/json")

	updateBook := &storage.Book{}
	utils.ParseBody(req, updateBook)

	params := mux.Vars(req)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Printf("error while parsing ~ %s", err)
	}

	book, db := storage.GetBookById(ID)
	if updateBook.Name != "" {
		book.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		book.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication
	}

	db.Save(&book)
	response, _ := json.Marshal(book)
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}
