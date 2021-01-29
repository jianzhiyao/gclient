package gclient

import (
	"github.com/jianzhiyao/gclient/request"
)

//NewRequest new request
func NewRequest(method, url string) (*request.Request, error) {
	return request.New(method, url)
}

//NewRequestGet new request of http.MethodGet
func NewRequestGet(url string) (*request.Request, error) {
	return request.Get(url)
}

//NewRequestHead new request of http.MethodHead
func NewRequestHead(url string) (*request.Request, error) {
	return request.Head(url)
}

//NewRequestPost new request of http.MethodPost
func NewRequestPost(url string) (*request.Request, error) {
	return request.Post(url)
}

//NewRequestPut new request of http.MethodPut
func NewRequestPut(url string) (*request.Request, error) {
	return request.Put(url)
}

//NewRequestPatch new request of http.MethodPatch
func NewRequestPatch(url string) (*request.Request, error) {
	return request.Patch(url)
}

//NewRequestDelete new request of http.MethodDelete
func NewRequestDelete(url string) (*request.Request, error) {
	return request.Delete(url)
}

//NewRequestConnect new request of http.MethodConnect
func NewRequestConnect(url string) (*request.Request, error) {
	return request.Connect(url)
}

//NewRequestOptions new request of http.MethodOptions
func NewRequestOptions(url string) (*request.Request, error) {
	return request.Options(url)
}

//NewRequestTrace new request of http.MethodTrace
func NewRequestTrace(url string) (*request.Request, error) {
	return request.Trace(url)
}
