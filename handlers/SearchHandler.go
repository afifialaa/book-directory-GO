package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	database "github.com/afifialaa/REST-GO/database"

	"github.com/fatih/structs"
)

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
	id := req.FormValue("id")
	result := database.SearchByID(id)
	json.NewEncoder(w).Encode(result)
}

func TestHandler(res http.ResponseWriter, req *http.Request){
	title := req.FormValue("title")
	fmt.Println("title in handler is " , title)
	result := database.SearchByTitle(title)
	data := structs.Map(result)
	json.NewEncoder(res).Encode(data)
}
