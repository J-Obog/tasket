package api

type RestRequest struct {
	Url         string
	UrlParams   map[string]interface{}
	QueryParams map[string]interface{}
	Body        []byte
}

type RestResponse struct {
	Object interface{}
	Status int
}
