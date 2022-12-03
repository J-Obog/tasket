package rest

type RestHandler func(req RestRequest) (res RestResponse, status int)

type RestRequest struct {
	Url         string
	UrlParams   map[string]interface{}
	QueryParams map[string]interface{}
	Metadata    map[string]interface{}
	Body        []byte
}

type RestError struct {
	Type   string `json:"type"`
	Detail string `json:"detail"`
}

type RestResponse struct {
	Data    interface{} `json:"data"`
	Error   RestError   `json:"error"`
	Message string      `json:"message"`
}

func ServerErrorResponse() RestResponse {
	return RestResponse{
		Error: RestError{
			Type:   "server error",
			Detail: "internal error occured while processing request",
		},
	}
}

func UserId(requestMeta map[string]interface{}) string {
	return requestMeta["userId"].(string)
}
