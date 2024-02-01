package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/winhandsome309/gobook/pkg/models"
	"github.com/winhandsome309/gobook/pkg/utils"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is Main Page"))

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	bookToCreate := &models.Book{}
	utils.ParseBody(r, bookToCreate)
	book := bookToCreate.CreateBook()
	res, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBook()
	res, _ := json.MarshalIndent(books, "", "	")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		panic(err)
	}
	b := models.GetBookById(int(id))
	res, _ := json.MarshalIndent(b, "", "	")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId, err := strconv.ParseInt(vars["bookId"], 10, 64)
	if err != nil {
		panic(err)
	}
	bookNew := &models.Book{}
	utils.ParseBody(r, bookNew)
	bookUpdated := models.GetBookById(int(bookId))
	bookUpdated.UpdateBook(bookNew)
	w.WriteHeader(http.StatusOK)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId, err := strconv.ParseInt(vars["bookId"], 10, 64)
	if err != nil {
		panic(err)
	}
	models.DeleteBook(int(bookId))
	w.WriteHeader(http.StatusOK)
}
