package models

type ResponseError struct {
	Error interface{} `json:"error"`
}

type ServerError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}
