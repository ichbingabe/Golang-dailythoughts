package controller

import (
	"fmt"
	"log"
	"net/http"

	controller "com.webapp/webapp/controller/app"
	"github.com/gorilla/mux"
)

var router *mux.Router

func initHandlers() {
	router.HandleFunc("/api/thoughts", controller.GetAllThoughts).Methods("GET")
	router.HandleFunc("/", controller.SayHello).Methods("GET")
	router.HandleFunc("/api/thought/{id}", controller.GetThoughtById).Methods("GET")
}

func Start() {
	router = mux.NewRouter()

	initHandlers()
	fmt.Printf("router initialed and listening on 3200\n")
	log.Fatal(http.ListenAndServe(":3200", router))
}
