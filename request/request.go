package request

import (
	"errors"
	"net/http"
)

type Request struct {
	method  string
	url     string
	headers map[string]string
}

func New(method, url string) (*Request, error) {
	switch method {
	case http.MethodGet:
	case http.MethodPost:
	case http.MethodConnect:
	case http.MethodDelete:
	case http.MethodHead:
	case http.MethodOptions:
	case http.MethodPatch:
	case http.MethodPut:
	case http.MethodTrace:
	default:
		return nil, errors.New("not a valid http method")
	}
	return &Request{
		method:  method,
		url:     url,
		headers: make(map[string]string),
	}, nil
}

func (r *Request) SetHeader(key string, value string) {
	r.headers[key] = value
}

func (r *Request) Method() string {
	return r.method
}

func (r *Request) Headers() map[string]string {
	return r.headers
}

