package connect

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"	
	"log"
)

var connection *gorm.DB

const engine_sql string = "mysql"

const username string = "root"
const password string = "ArqSoft.123"
const database string = "ChatsDB"

func InitializeDatabase(){
	connection = ConnectORM(CreateString())
	log.Println("Conexión exitosa")
}

func CloseConnection(){
	connection.Close()
	log.Println("Conexión cerrada")
}

func ConnectORM(stringConnection string) *gorm.DB{
	connection, err := gorm.Open(engine_sql, stringConnection)
	if(err != nil){
		log.Fatal(err)
		return nil
	}
	return connection
}

func CreateString() string{
	return username+":"+password+"@/"+database
}