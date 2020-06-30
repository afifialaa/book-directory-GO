package handlers

import (
	"encoding/json"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	res := map[string]string{"msg": "Hello from home page"}
	json.NewEncoder(w).Encode(res)
}
