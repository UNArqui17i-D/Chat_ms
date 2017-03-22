package structures

import (
	"gopkg.in/mgo.v2/bson"
)

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
