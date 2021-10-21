package gclient

import (
	"github.com/jianzhiyao/gclient/tests"
	"net/http"
	"testing"
)

func TestNewRequest(t *testing.T) {
	if c, err := NewRequest(http.MethodPost, tests.GetServerUrl()); c == nil || err != nil {
		t.Error()
	}
}

func TestNewRequestDelete(t *testing.T) {
	if c, _ := NewRequestDelete(tests.GetServerUrl()); c.GetMethod() != http.MethodDelete {
		t.Error()
	}
}

func TestNewRequestGet(t *testing.T) {
	if c, _ := NewRequestGet(tests.GetServerUrl()); c.GetMethod() != http.MethodGet {
		t.Error()
	}
}

func TestNewRequestHead(t *testing.T) {
	if c, _ := NewRequestHead(tests.GetServerUrl()); c.GetMethod() != http.MethodHead {
		t.Error()
	}
}

func TestNewRequestPatch(t *testing.T) {
	if c, _ := NewRequestPatch(tests.GetServerUrl()); c.GetMethod() != http.MethodPatch {
		t.Error()
	}
}

func TestNewRequestTrace(t *testing.T) {
	if c, _ := NewRequestTrace(tests.GetServerUrl()); c.GetMethod() != http.MethodTrace {
		t.Error()
	}
}

func TestNewRequestPost(t *testing.T) {
	if c, _ := NewRequestPost(tests.GetServerUrl()); c.GetMethod() != http.MethodPost {
		t.Error(c.GetMethod())
	}
}

func TestNewRequestPut(t *testing.T) {
	if c, _ := NewRequestPut(tests.GetServerUrl()); c.GetMethod() != http.MethodPut {
		t.Error()
	}
}

func TestNewRequestConnect(t *testing.T) {
	if c, _ := NewRequestConnect(tests.GetServerUrl()); c.GetMethod() != http.MethodConnect {
		t.Error()
	}
}

func TestNewRequestOptions(t *testing.T) {
	if c, _ := NewRequestOptions(tests.GetServerUrl()); c.GetMethod() != http.MethodOptions {
		t.Error()
	}
}
