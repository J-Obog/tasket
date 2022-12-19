package resource

import (
	"net/http"

	"github.com/J-Obog/tasket/src/rest"
)

func MakeResponse(data interface{}, status int64) rest.RestResponse {
	return rest.RestResponse{
		Data:   data,
		Status: status,
	}
}

func MakeError(errType string, detail string) rest.RestError {
	return rest.RestError{
		Type:   errType,
		Detail: detail,
	}
}

func MakeMessage(message string) rest.RestMessage {
	return rest.RestMessage{
		Message: message,
	}
}

func ServerErrorResponse() rest.RestResponse {
	resErr := MakeError("server error", "internal error occured while processing request")
	return MakeResponse(&resErr, http.StatusOK)
}

func UserId(requestMeta map[string]interface{}) string {
	return requestMeta["userId"].(string)
}
