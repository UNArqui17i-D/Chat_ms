package connect

import (
	"log"

	"../structures"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//var connection *gorm.DB

var (
	session    *mgo.Session
	collection *mgo.Collection
)

const engine_sql string = "mysql"

const username string = "root"
const password string = "ArqSoft.123"
const database string = "ChatsDB"

func InitializeDatabase() {
	session, err := mgo.Dial("localhost")
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
func GetMessage(id bson.ObjectId) (structures.Message, error) {
	chat := structures.Message{}
	err := collection.Find(bson.M{"_id": id}).One(&chat)
	return chat, err
}

func GetChat(userFrom string, userTo string) ([]structures.Message, error) {
	chats := []structures.Message{}
	log.Println(userFrom)
	err := collection.Find(bson.M{"userfrom": userFrom}).All(&chats)
	return chats, err
}

//Funciones POST
func CreateMessage(message structures.Message) (structures.Message, error) {
	obj_id := bson.NewObjectId()
	message.Id = obj_id
	err := collection.Insert(&message)
	return message, err
}
