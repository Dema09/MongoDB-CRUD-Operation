package response

import "net/http"

const StatusOK = "OK"

type RestBody struct{
	StatusCode int `json:"status_code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func NewInternalServerError(message string) *RestBody {
	return &RestBody{
		StatusCode: http.StatusInternalServerError,
		Message: message,
		Data: nil,
	}
}

func NewBadRequest(message string) *RestBody{
	return &RestBody{
		StatusCode: http.StatusBadRequest,
		Message: message,
		Data: nil,
	}
}

func NewStatusOK(data interface{}) *RestBody{
	return &RestBody{
		StatusCode: http.StatusOK,
		Message: StatusOK,
		Data: data,
	}
}

func NewStatusCreated(message string, data interface{}) *RestBody{
	return &RestBody{
		StatusCode: http.StatusCreated,
		Message: message,
		Data: data,
	}
}
