package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitServer() {
	userRouter := mux.NewRouter().PathPrefix("/api/user").Subrouter()
	userRouter.HandleFunc("", HandleGetUser).Methods("GET")
	userRouter.HandleFunc("", HandleCreateUser).Methods("POST")

	http.Handle("/api/user", userRouter)
}

func Serve() {
	log.Default().Println("Starting server on port 5001...")
	err := http.ListenAndServe(":5001", nil)
	if err != nil {
		log.Default().Fatal(err)
	}
}
