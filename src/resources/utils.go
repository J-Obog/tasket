package resources

import (
	"net/http"

	"github.com/J-Obog/tasket/src/types"
)

func MakeResponse(data interface{}, status int64) types.RestResponse {
	return types.RestResponse{
		Data:   data,
		Status: status,
	}
}

func MakeError(errType string, detail string) types.RestError {
	return types.RestError{
		Type:   errType,
		Detail: detail,
	}
}

func MakeMessage(message string) types.RestMessage {
	return types.RestMessage{
		Message: message,
	}
}

func ServerErrorResponse() types.RestResponse {
	resErr := MakeError("server error", "internal error occured while processing request")
	return MakeResponse(&resErr, http.StatusOK)
}

func UserId(requestMeta map[string]interface{}) string {
	return requestMeta["userId"].(string)
}
