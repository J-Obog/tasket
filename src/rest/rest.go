package rest

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

type RestMessage struct {
	Message string `json:"message"`
}

type RestResponse struct {
	Data   interface{}
	Status int64
}
