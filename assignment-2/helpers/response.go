package helpers

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(message string, data interface{}) Response {
	res := Response{
		Status:  true,
		Message: message,
		Data:    data,
	}
	return res
}
