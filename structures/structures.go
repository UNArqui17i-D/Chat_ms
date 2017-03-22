package structures

type Message struct {
	Id       int    `json:"id"`
	Message  string `json:"message"`
	Userfrom int    `json:"userfrom"`
	Userto   int    `json:"userto"`
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
