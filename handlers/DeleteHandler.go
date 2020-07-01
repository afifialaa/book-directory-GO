package handlers

import (
	"encoding/json"
	"net/http"

	database "github.com/afifialaa/REST-GO/database"
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
