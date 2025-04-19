package serializer

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
	Count   int64       `json:"count"`
}

type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}
