package request

import "net/http"

//Get generate a new request of http.MethodGet
func Get(url string) (*Request, error) {
	return New(http.MethodGet, url)
}

//Head generate a new request of http.MethodHead
func Head(url string) (*Request, error) {
	return New(http.MethodHead, url)
}

//Post generate a new request of http.MethodPost
func Post(url string) (*Request, error) {
	return New(http.MethodPost, url)
}

//Put generate a new request of http.MethodPut
func Put(url string) (*Request, error) {
	return New(http.MethodPut, url)
}

//Patch generate a new request of http.MethodPatch
func Patch(url string) (*Request, error) {
	return New(http.MethodPatch, url)
}

//Delete generate a new request of http.MethodDelete
func Delete(url string) (*Request, error) {
	return New(http.MethodDelete, url)
}

//Connect generate a new request of http.MethodConnect
func Connect(url string) (*Request, error) {
	return New(http.MethodConnect, url)
}

//Options generate a new request of http.MethodOptions
func Options(url string) (*Request, error) {
	return New(http.MethodOptions, url)
}

//Trace generate a new request of http.MethodTrace
func Trace(url string) (*Request, error) {
	return New(http.MethodTrace, url)
}
