package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
	"./connects"
)

//Estructuras
type Response struct{
	Message string  `json:"message"`
	Status bool `json:"status"`
}
type Message struct{
	Message string  `json:"message"`
	IdUserFrom int `json:"from"`
	IdUserTo int `json:"to"`
}

//Función Main
func main(){
	connect.InitializeDatabase()
	defer connect.CloseConnection()
	mux := mux.NewRouter()
	mux.HandleFunc("/Chat_ms/Api/Message/{userId}", GetMessages).Methods("GET")
	mux.HandleFunc("/HolaJson", HolaMundoJson).Methods("GET")
	mux.HandleFunc("/Chat_ms/Api/Message", AddMessage).Methods("POST")

	http.Handle("/", mux)
	log.Println("El servidor encuentra en el puerto 4005")
	log.Fatal(http.ListenAndServe(":4005", nil))
}


//Funciones GET
func GetMessages(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	userId := vars["userId"]
	w.Write([]byte(userId))
}

func HolaMundoJson(w http.ResponseWriter, r *http.Request){
	Response := CreateResponse()
	json.NewEncoder(w).Encode(Response)
}

func CreateResponse() Response {
	return Response{"Esto está en formato JSON", true}
}


//Funciones POST
func AddMessage(w http.ResponseWriter, r *http.Request) {
	//Message := Message{"Probando", 1, 2}
}