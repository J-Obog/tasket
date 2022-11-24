package models

type RestRequest struct {
	Url         string
	UrlParams   map[string]interface{}
	QueryParams map[string]interface{}
	Body        []byte
}

type RestResponse struct {
	Body   []byte
	Status int
}
