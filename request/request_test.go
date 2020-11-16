package request

import (
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	methods := []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodConnect,
		http.MethodDelete,
		http.MethodHead,
		http.MethodOptions,
		http.MethodPatch,
		http.MethodPut,
		http.MethodTrace,
	}
	url := "https://cn.bing.com"
	for _, method := range methods {
		req, err := New(method, url)
		if err != nil {
			t.Error(err)
			return
		}

		if req.method != method {
			t.Error(err)
			return
		}

		if req.url != url {
			t.Error(err)
			return
		}
	}

}

func TestNew2(t *testing.T) {
	method, url := `1`, "https://cn.bing.com"
	_, err := New(method, url)
	if err == nil {
		t.Error(err)
		return
	}
}

func TestRequest_Method(t *testing.T) {
	method, url := http.MethodPatch, "https://cn.bing.com"
	req, _ := New(method, url)

	if req.Method() != method {
		t.Error()
		return
	}

	if req.Method() == http.MethodGet {
		t.Error()
		return
	}
}

func TestRequest_Headers(t *testing.T) {
	method, url := http.MethodPatch, "https://cn.bing.com"
	req, _ := New(method, url)

	req.SetHeader("a", "1")
	req.SetHeader("a", "1")
	req.SetHeader("b", "1")

	if len(req.Headers()) != 2 {
		t.Error()
		return
	}
}

func TestRequest_SetHeader(t *testing.T) {
	method, url := http.MethodPatch, "https://cn.bing.com"
	req, _ := New(method, url)

	req.SetHeader("a", "1")
	req.SetHeader("a", "1")
	req.SetHeader("b", "1")

	if _, ok := req.headers["a"]; !ok {
		t.Error()
		return
	}
	if _, ok := req.headers["b"]; !ok {
		t.Error()
		return
	}
	if _, ok := req.headers["c"]; ok {
		t.Error()
		return
	}
}

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