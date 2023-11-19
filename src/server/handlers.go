package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	db "tictactoe-service/database"
	"tictactoe-service/server/dto"
)

func HandleUser(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("Handling user...")
}

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	queryId := r.URL.Query().Get("id")
	log.Default().Println("Handling get user with id: " + queryId)

	id, err := strconv.ParseInt(queryId, 10, 64)
	if err != nil {
		log.Default().Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := db.GetUser(id)
	if err != nil {
		log.Default().Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	jsonBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Default().Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	body := &dto.CreateUserRequest{}
	json.Unmarshal(jsonBody, body)
	if body.Username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := db.CreateUser(body.Username)
	if err != nil {
		log.Default().Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}