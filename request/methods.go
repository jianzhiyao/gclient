package request

import "net/http"

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
