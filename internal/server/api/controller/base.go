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

func response404(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(message))
}

func response400(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

func response500(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(message))
}
