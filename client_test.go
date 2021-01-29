package gclient

import (
	"net/http"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	if c := New(); c == nil {
		t.Error()
	}
}

func TestClient_Option(t *testing.T) {
	timeout := 5423 * time.Second
	c := New()
	c.Option(OptTimeout(timeout))

	if c.clientTimeout != timeout {
		t.Error()
	}
}

func TestClient_Options(t *testing.T) {
	timeout := 12398 * time.Second
	retry := 86543

	c := New()
	c.Options(
		OptTimeout(timeout),
		OptRetry(retry),
	)

	if c.clientTimeout != timeout {
		t.Error()
	}
	if c.retry != retry {
		t.Error()
	}
}

func TestClient_Do(t *testing.T) {
	c := New()
	url := `https://cn.bing.com`

	if resp, err := c.Do(http.MethodGet, url); err != nil {
		t.Error(err)
	} else if resp.StatusCode() != http.StatusOK {
		t.Error()
	}
}

func TestClient_DoRequest(t *testing.T) {
	c := New()
	url := `https://cn.bing.com`

	req, _ := NewRequest(http.MethodGet, url)
	if resp, err := c.DoRequest(req); err != nil {
		t.Error(err)
	} else if resp.StatusCode() != http.StatusOK {
		t.Error()
	}
}