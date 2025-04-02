package serializer

// 一个通用的基础序列化器
type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Count   int64       `json:"count"`
}

type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}
