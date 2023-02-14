package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"com.webapp/webapp/model"
	"github.com/gorilla/mux"
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

func GetThoughtById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	thought, err := model.GetThoughtById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(thought)
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello World!"))
}
