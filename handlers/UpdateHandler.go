package handlers

import (
	"encoding/json"
	"net/http"

	database "github.com/afifialaa/REST-GO/database"
	"github.com/afifialaa/REST-GO/models"
)

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
