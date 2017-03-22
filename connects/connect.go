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

	//connection = ConnectORM(CreateString())
	log.Println("Conexión exitosa")
}

func CloseConnection() {
	defer session.Close()
	//connection.Close()
	log.Println("Conexión cerrada")
}

/*func ConnectORM(stringConnection string) *gorm.DB {
	connection, err := gorm.Open(engine_sql, stringConnection)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return connection
}

func CreateString() string {
	return username + ":" + password + "@/" + database
}*/

//Funciones GET
func GetMessage(id string) structures.Message {
	chat := structures.Message{}
	collection.Find(bson.M{"id": id})
	//connection.Where("id = ?", id).Find(&chat)
	return chat
}

func GetChat(userFrom string, userTo string) []structures.Message {
	chats := []structures.Message{}
	log.Println(userFrom)
	//connection.Where("(userfrom = ? && userto = ?) || (userto = ? && userfrom = ?)", userFrom, userTo, userFrom, userTo).Find(&chats)
	return chats
}

//Funciones POST
func CreateMessage(message structures.Message) structures.Message {
	collection.Insert(&message)
	//connection.Create(&message)
	return message
}
