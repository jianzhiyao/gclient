package request

import (
	"net/http"
	"testing"
)

func TestConnect(t *testing.T) {
	url := "https://cn.bing.com"
	req, _ := Connect(url)

	if req.Method() != http.MethodConnect {
		t.Error()
		return
	}
}

func TestDelete(t *testing.T) {
	url := "https://cn.bing.com"
	req, _ := Delete(url)

	if req.Method() != http.MethodDelete {
		t.Error()
		return
	}
}

func TestGet(t *testing.T) {
	url := "https://cn.bing.com"
	req, _ := Get(url)

	if req.Method() != http.MethodGet {
		t.Error()
		return
	}
}

func TestHead(t *testing.T) {
	url := "https://cn.bing.com"
	req, _ := Head(url)

	if req.Method() != http.MethodHead {
		t.Error()
		return
	}
}

func TestOptions(t *testing.T) {
	url := "https://cn.bing.com"
	req, _ := Options(url)

	if req.Method() != http.MethodOptions {
		t.Error()
		return
	}
}

func TestPatch(t *testing.T) {
	url := "https://cn.bing.com"
	req, _ := Patch(url)

	if req.Method() != http.MethodPatch {
		t.Error()
		return
	}
}

func TestPost(t *testing.T) {
	url := "https://cn.bing.com"
	req, _ := Post(url)

	if req.Method() != http.MethodPost {
		t.Error()
		return
	}
}

func TestPut(t *testing.T) {
	url := "https://cn.bing.com"
	req, _ := Put(url)

	if req.Method() != http.MethodPut {
		t.Error()
		return
	}
}

func TestTrace(t *testing.T) {
	url := "https://cn.bing.com"
	req, _ := Trace(url)

	if req.Method() != http.MethodTrace {
		t.Error()
		return
	}
}
