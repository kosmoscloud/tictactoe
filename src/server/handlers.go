package server

import (
	"log"
	"net/http"
)

func HandleUser(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("Handling user...")
}
