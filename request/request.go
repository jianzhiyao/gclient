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
		return nil, errors.New("Not a valid http method")
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


func Get(url string) (*Request, error) {
	return New(http.MethodGet, url)
}

func Head(url string) (*Request, error) {
	return New(http.MethodHead, url)
}

func Post(url string) (*Request, error) {
	return New(http.MethodPost, url)
}

func Put(url string) (*Request, error) {
	return New(http.MethodPut, url)
}

func Patch(url string) (*Request, error) {
	return New(http.MethodPatch, url)
}

func Delete(url string) (*Request, error) {
	return New(http.MethodDelete, url)
}

func Connect(url string) (*Request, error) {
	return New(http.MethodConnect, url)
}

func Options(url string) (*Request, error) {
	return New(http.MethodOptions, url)
}

func Trace(url string) (*Request, error) {
	return New(http.MethodTrace, url)
}