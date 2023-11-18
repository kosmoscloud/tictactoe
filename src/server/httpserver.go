package server

import (
	"log"
	"net/http"
)

func InitServer() {
	http.HandleFunc("/api/user", HandleUser)
}

func Serve() {
	log.Default().Println("Starting server on port 5001...")
	err := http.ListenAndServe(":5001", nil)
	if err != nil {
		log.Default().Fatal(err)
	}
}
