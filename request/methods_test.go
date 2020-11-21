package request

import (
	"net/http"
	"testing"
)

func TestConnect(t *testing.T) {
	url := "https://example.com/?for=TestConnect"
	req, _ := Connect(url)

	if req.Method() != http.MethodConnect {
		t.Error()
		return
	}

	if req.Url() != url {
		t.Error()
		return
	}
}

func TestDelete(t *testing.T) {
	url := "https://example.com/?for=TestDelete"
	req, _ := Delete(url)

	if req.Method() != http.MethodDelete {
		t.Error()
		return
	}

	if req.Url() != url {
		t.Error()
		return
	}
}

func TestGet(t *testing.T) {
	url := "https://example.com/?for=TestGet"
	req, _ := Get(url)

	if req.Method() != http.MethodGet {
		t.Error()
		return
	}

	if req.Url() != url {
		t.Error()
		return
	}
}

func TestHead(t *testing.T) {
	url := "https://example.com/?for=TestHead"
	req, _ := Head(url)

	if req.Method() != http.MethodHead {
		t.Error()
		return
	}

	if req.Url() != url {
		t.Error()
		return
	}
}

func TestOptions(t *testing.T) {
	url := "https://example.com/?for=TestOptions"
	req, _ := Options(url)

	if req.Method() != http.MethodOptions {
		t.Error()
		return
	}

	if req.Url() != url {
		t.Error()
		return
	}
}

func TestPatch(t *testing.T) {
	url := "https://example.com/?for=TestPatch"
	req, _ := Patch(url)

	if req.Method() != http.MethodPatch {
		t.Error()
		return
	}

	if req.Url() != url {
		t.Error()
		return
	}
}

func TestPost(t *testing.T) {
	url := "https://example.com/?for=TestPost"
	req, _ := Post(url)

	if req.Method() != http.MethodPost {
		t.Error()
		return
	}

	if req.Url() != url {
		t.Error()
		return
	}
}

func TestPut(t *testing.T) {
	url := "https://example.com/?for=TestPut"
	req, _ := Put(url)

	if req.Method() != http.MethodPut {
		t.Error()
		return
	}

	if req.Url() != url {
		t.Error()
		return
	}
}

func TestTrace(t *testing.T) {
	url := "https://example.com/?for=TestTrace"
	req, _ := Trace(url)

	if req.Method() != http.MethodTrace {
		t.Error()
		return
	}

	if req.Url() != url {
		t.Error()
		return
	}
}
