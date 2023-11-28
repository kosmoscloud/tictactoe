package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitServer() {
	CuserRouter := mux.NewRouter().PathPrefix("/api/user").Subrouter()
	CuserRouter.HandleFunc("", HandleCreateUser).Methods("POST")

	RUDuserRouter := mux.NewRouter().PathPrefix("/api/user").Subrouter()
	RUDuserRouter.HandleFunc("/{id}", HandleGetUser).Methods("GET")
	RUDuserRouter.HandleFunc("/{id}", HandleUpdateUser).Methods("PUT")
	RUDuserRouter.HandleFunc("/{id}", HandleDeleteUser).Methods("DELETE")

	http.Handle("/api/user", CuserRouter)
	http.Handle("/api/user/", RUDuserRouter)

	CroomRouter := mux.NewRouter().PathPrefix("/api/room").Subrouter()
	CroomRouter.HandleFunc("", HandleCreateRoom).Methods("POST")

	RUDroomRouter := mux.NewRouter().PathPrefix("/api/room").Subrouter()
	//RUDroomRouter.HandleFunc("/{id}", HandleGetRoom).Methods("GET")
	//RUDroomRouter.HandleFunc("/{id}", HandleUpdateRoom).Methods("PUT")
	RUDroomRouter.HandleFunc("/{id}", HandleDeleteRoom).Methods("DELETE")

	http.Handle("/api/room", CroomRouter)
	http.Handle("/api/room/", RUDroomRouter)

}

func Serve() {
	log.Default().Println("Starting server on port 5001...")
	err := http.ListenAndServe(":5001", nil)
	if err != nil {
		log.Default().Fatal(err)
	}
}
