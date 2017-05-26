package main

import (
	"encoding/json"
	"log"
	"net/http"

	"strconv"

	//"./connects"
	//"./structures"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Structures

type Message struct {
	Id       bson.ObjectId `bson:"_id" json:"id"`
	Message  string        `bson:"message" json:"message"`
	Userfrom int           `bson:"userfrom" json:"userfrom"`
	Userto   int           `bson:"userto" json:"userto"`
}

type Response struct {
	Status  string    `json:"status"`
	Data    []Message `json:"data"`
	Message string    `json:"message"`
}

type ResponseMessage struct {
	Status  string  `json:"status"`
	Data    Message `json:"data"`
	Message string  `json:"message"`
}






//Función Main
func main() {
	InitializeDatabase()
	defer CloseConnection()
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
	chat, err := DBGetMessage(id)

	if err != nil {
		status = "Error"
		message = "No existe este chat"
	}
	response := ResponseMessage{status, chat, message}

	json.NewEncoder(w).Encode(response)
}

func GetChat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idFrom, _ := strconv.Atoi(vars["userFrom"])
	idTo, _ := strconv.Atoi(vars["userTo"])

	status := "success"
	var message string
	chat, err := DBGetChat(idFrom, idTo)

	if err != nil {
		status = "Error"
		message = "No existe este chat"
	}
	response := Response{status, chat, message}
	json.NewEncoder(w).Encode(response)
}

//Funciones POST
func AddMessage(w http.ResponseWriter, r *http.Request) {
	message, err := DBCreateMessage(GetMessageRequest(r))
	response := ResponseMessage{}
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
func GetMessageRequest(r *http.Request) Message {
	message := Message{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&message)
	if err != nil {
		log.Fatal(err)
	}
	return message
}





//Connect

var (
	session    *mgo.Session
	collection *mgo.Collection
)

const engine_sql string = "mysql"

const username string = "root"
const password string = "ArqSoft.123"
const database string = "ChatsDB"

func InitializeDatabase() {
	session, err := mgo.Dial("192.168.99.101:3307")
	if err != nil {
		log.Fatal(err)
	}
	session.SetMode(mgo.Monotonic, true)
	collection = session.DB("MessagesDB").C("messages")

	log.Println("Conexión exitosa")
}

func CloseConnection() {
	defer session.Close()
	log.Println("Conexión cerrada")
}

//Funciones GET
func DBGetMessage(id bson.ObjectId) (Message, error) {
	chat := Message{}
	err := collection.Find(bson.M{"_id": id}).One(&chat)
	return chat, err
}

func DBGetChat(userFrom int, userTo int) ([]Message, error) {
	chats := []Message{}
	log.Println(userFrom)
	err := collection.Find(bson.M{"$or": []bson.M{bson.M{"userfrom": userFrom, "userto": userTo}, bson.M{"userfrom": userTo, "userto": userFrom}}}).All(&chats)
	return chats, err
}

//Funciones POST
func DBCreateMessage(message Message) (Message, error) {
	obj_id := bson.NewObjectId()
	message.Id = obj_id
	err := collection.Insert(&message)
	return message, err
}

