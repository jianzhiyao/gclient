package gclient

import (
	"net/http"
	"testing"
)

func TestNewRequest(t *testing.T) {
	if c, err := NewRequest(http.MethodPost, ``); c == nil || err != nil {
		t.Error()
	}
}

func TestNewRequestDelete(t *testing.T) {
	if c, _ := NewRequestDelete(``); c.GetMethod() != http.MethodDelete {
		t.Error()
	}
}

func TestNewRequestGet(t *testing.T) {
	if c, _ := NewRequestGet(``); c.GetMethod() != http.MethodGet {
		t.Error()
	}
}

func TestNewRequestHead(t *testing.T) {
	if c, _ := NewRequestHead(``); c.GetMethod() != http.MethodHead {
		t.Error()
	}
}

func TestNewRequestPatch(t *testing.T) {
	if c, _ := NewRequestPatch(``); c.GetMethod() != http.MethodPatch {
		t.Error()
	}
}

func TestNewRequestTrace(t *testing.T) {
	if c, _ := NewRequestTrace(``); c.GetMethod() != http.MethodTrace {
		t.Error()
	}
}

func TestNewRequestPost(t *testing.T) {
	if c, _ := NewRequestPost(``); c.GetMethod() != http.MethodPost {
		t.Error()
	}
}

func TestNewRequestPut(t *testing.T) {
	if c, _ := NewRequestPut(``); c.GetMethod() != http.MethodPut {
		t.Error()
	}
}

func TestNewRequestConnect(t *testing.T) {
	if c, _ := NewRequestConnect(``); c.GetMethod() != http.MethodConnect {
		t.Error()
	}
}

func TestNewRequestOptions(t *testing.T) {
	if c, _ := NewRequestOptions(``); c.GetMethod() != http.MethodOptions {
		t.Error()
	}
}
