package handlers

import (
	"encoding/json"
	"net/http"

	database "book-directory-GO/database"
	"book-directory-GO/models"

	"github.com/fatih/structs"
	"github.com/gorilla/mux"
)

func DeleteByID(w http.ResponseWriter, req *http.Request) {
	bookId := req.FormValue("id")
	ans := database.DeleteByID(bookId)

	if !ans {
		data := map[string]string{"msg": "failed to delete"}
		json.NewEncoder(w).Encode(data)
	} else {
		data := map[string]string{"msg": "book was deleted"}
		json.NewEncoder(w).Encode(data)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	res := map[string]string{"msg": "Hello from home page"}
	json.NewEncoder(w).Encode(res)
}

func UpdateBook(w http.ResponseWriter, req *http.Request) {

	var book = models.BookType{
		BookID:             req.FormValue("id"),
		Title:              req.FormValue("title"),
		Authors:            req.FormValue("authors"),
		Average_rating:     req.FormValue("averageRating"),
		Isbn:               req.FormValue("isbn"),
		Isbn13:             req.FormValue("isbn13"),
		Language_code:      req.FormValue("languageCode"),
		Ratings_count:      req.FormValue("ratingsCount"),
		Text_reviews_count: req.FormValue("textReviewCount"),
		Publication_date:   req.FormValue("publicationDate"),
		Publisher:          req.FormValue("publisher"),
	}

	result := database.UpdateBook(book)

	if !result {
		json.NewEncoder(w).Encode(map[string]string{"msg": "failed to update book"})
	}
	json.NewEncoder(w).Encode(map[string]string{"msg": "book was updated"})

}

func SearchByTitle(w http.ResponseWriter, req *http.Request) {
	title := req.FormValue("title")
	result := database.SearchByTitle(title)
	res := structs.Map(result)
	json.NewEncoder(w).Encode(res)
}

func SearchByAuthor(w http.ResponseWriter, req *http.Request) {
	author := req.FormValue("author")
	result := database.SearchByAuthor(author)

	json.NewEncoder(w).Encode(result)
}

func SearchByID(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["bookId"]
	result := database.SearchByID(id)
	json.NewEncoder(w).Encode(result)
}
