package structures

type Chat struct{
	Id 			int 	`json:"id"`
	Message 	string  `json:"message"`
	UserFrom 	int 	`json:"idUserFrom"`
	UserTo 		int 	`json:"idUserTo"`
}


type Response struct{
	Status 		string 	`json:"status"`
	Data 		Chat 	`json:"data"`
	Message		string 	`json:"message"`
}