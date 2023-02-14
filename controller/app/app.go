package controller

import (
	"encoding/json"
	"net/http"

	"com.webapp/webapp/model"
)

func GetAllThoughts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	thoughts, err := model.GetAllThoughts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(thoughts)
	}
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello World!"))
}
