package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	db "tictactoe-service/database"
	"tictactoe-service/server/dto"

	"github.com/gorilla/mux"
)

func HandleUser(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("Handling user...")
}

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	sid := pathParams["id"]
	log.Default().Println("Handling update user with id: " + sid)

	id, err := strconv.ParseInt(sid, 10, 64)
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

func HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	sid := pathParams["id"]
	log.Default().Println("Handling update user with id: " + sid)

	id, err := strconv.ParseInt(sid, 10, 64)
	jsonBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Default().Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body := &dto.UpdateUserRequest{UserId: id}
	json.Unmarshal(jsonBody, body)

	if body.Username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := db.UpdateUser(body.UserId, body.Username)
	if err != nil {
		log.Default().Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Default().Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	id := pathParams["id"]
	log.Default().Println("Handling delete user with id: " + id)

	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Default().Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := db.DeleteUser(userId)
	if err != nil {
		log.Default().Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Default().Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
