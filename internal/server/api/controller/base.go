package controller

import (
	"encoding/json"
	"net/http"
)

func responseJSON(w http.ResponseWriter, v interface{}) {
	res, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
