package utils

import (
	"net/http"

	"github.com/J-Obog/tasket/src/api"
)

func GenerateUUID() string {
	return ""
}

func TimeNow() int64 {
	return 0
}

func GenerateHash(s string) []byte {
	return nil
}

func CurrentUserId(req api.RestRequest) string {
	return ""
}

func MakeServerError() api.RestResponse {
	return api.RestResponse{
		Object: map[string]interface{}{
			"message": "Internal Server Error",
		},
		Status: http.StatusInternalServerError,
	}
}

func TaskIdParam() string {
	return ""
}
