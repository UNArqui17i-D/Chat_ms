package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HolaMundo(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hola mundo"))
}

func main(){
	mux := mux.NewRouter()
	mux.HandleFunc("/Hola", HolaMundo).Methods("GET")

	http.Handle("/", mux)
	log.Println("El servidor encuentra en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}