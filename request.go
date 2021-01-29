package gclient

import (
	"github.com/jianzhiyao/gclient/request"
)

//NewRequest new request
func NewRequest(method, url string) (*request.Request, error) {
	return request.New(method, url)
}

//NewRequestGet new request of http.Get
func NewRequestGet(url string) (*request.Request, error) {
	return request.Get(url)
}

//NewRequestHead new request of http.Head
func NewRequestHead(url string) (*request.Request, error) {
	return request.Head(url)
}

//NewRequestPost new request of http.Post
func NewRequestPost(url string) (*request.Request, error) {
	return request.Post(url)
}

//NewRequestPut new request of http.Put
func NewRequestPut(url string) (*request.Request, error) {
	return request.Put(url)
}

//NewRequestPatch new request of http.Patch
func NewRequestPatch(url string) (*request.Request, error) {
	return request.Patch(url)
}

//NewRequestDelete new request of http.Delete
func NewRequestDelete(url string) (*request.Request, error) {
	return request.Delete(url)
}
