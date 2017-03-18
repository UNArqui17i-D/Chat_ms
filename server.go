package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
)

type Response struct{
	Message string  `json:"message"`
	Status bool `json:"status"`
}

func HolaMundo(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hola mundo"))
}

func HolaMundoJson(w http.ResponseWriter, r *http.Request){
	Response := CreateResponse()
	json.NewEncoder(w).Encode(Response)
}

func CreateResponse() Response {
	return Response{"Esto está en formato JSON", true}
}

func main(){
	mux := mux.NewRouter()
	mux.HandleFunc("/Hola", HolaMundo).Methods("GET")
	mux.HandleFunc("/HolaJson", HolaMundoJson).Methods("GET")

	http.Handle("/", mux)
	log.Println("El servidor encuentra en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}