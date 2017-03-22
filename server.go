package main

import (
	"encoding/json"
	"log"
	"net/http"

	"strconv"

	"./connects"
	"./structures"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
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
	id := bson.ObjectIdHex(vars["id"])

	status := "success"
	var message string
	chat, err := connect.GetMessage(id)

	if err != nil {
		status = "Error"
		message = "No existe este chat"
	}
	response := structures.ResponseMessage{status, chat, message}

	json.NewEncoder(w).Encode(response)
}

func GetChat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idFrom, _ := strconv.Atoi(vars["userFrom"])
	idTo, _ := strconv.Atoi(vars["userTo"])

	status := "success"
	var message string
	chat, err := connect.GetChat(idFrom, idTo)

	if err != nil {
		status = "Error"
		message = "No existe este chat"
	}
	response := structures.Response{status, chat, message}
	json.NewEncoder(w).Encode(response)
}

//Funciones POST
func AddMessage(w http.ResponseWriter, r *http.Request) {
	message, err := connect.CreateMessage(GetMessageRequest(r))
	response := structures.ResponseMessage{}
	response.Data = message
	if err != nil {
		response.Status = "Error"
		response.Message = "No se pudo enviar el mensaje"
	} else {
		response.Status = "Success"
		response.Message = "Mensaje enviado"
	}
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
