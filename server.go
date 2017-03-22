package main

import (
	"encoding/json"
	"log"
	"net/http"

	"./connects"
	"./structures"
	"github.com/gorilla/mux"
)

//Función Main
func main() {
	connect.InitializeDatabase()
	defer connect.CloseConnection()
	mux := mux.NewRouter()
	mux.HandleFunc("/Chat_ms/Api/Message/{id}", GetMessage).Methods("GET")
	mux.HandleFunc("/Chat_ms/Api/Chat/{userFrom}&{userTo}", GetChat).Methods("GET")
	mux.HandleFunc("/Chat_ms/Api/Message", AddMessage).Methods("POST")

	http.Handle("/", mux)
	log.Println("El servidor encuentra en el puerto 4005")
	log.Fatal(http.ListenAndServe(":4005", nil))
}

//Funciones GET
func GetMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	status := "success"
	var message string
	chat := connect.GetMessage(id)

	if chat.Id <= 0 {
		status = "Error"
		message = "No existe este chat"
	}
	response := structures.ResponseMessage{status, chat, message}

	json.NewEncoder(w).Encode(response)
}

func GetChat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userFrom := vars["userFrom"]
	userTo := vars["userTo"]

	status := "success"
	var message string
	chat := connect.GetChat(userFrom, userTo)

	if len(chat) <= 0 {
		status = "Error"
		message = "No existe este chat"
	}
	response := structures.Response{status, chat, message}
	json.NewEncoder(w).Encode(response)
}

//Funciones POST
func AddMessage(w http.ResponseWriter, r *http.Request) {
	message := connect.CreateMessage(GetMessageRequest(r))
	response := structures.ResponseMessage{"success", message, "Mensaje enviado"}
	json.NewEncoder(w).Encode(response)
}

//Funciones privadas
func GetMessageRequest(r *http.Request) structures.Message {
	message := structures.Message{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&message)
	if err != nil {
		log.Fatal(err)
	}
	return message
}
