package gclient

import (
	"net/http"
	"os"
	"testing"
)

func TestNewRequest(t *testing.T) {
	if c, err := NewRequest(http.MethodPost, os.Getenv(`TEST_TARGET`)); c == nil || err != nil {
		t.Error()
	}
}

func TestNewRequestDelete(t *testing.T) {
	if c, _ := NewRequestDelete(os.Getenv(`TEST_TARGET`)); c.GetMethod() != http.MethodDelete {
		t.Error()
	}
}

func TestNewRequestGet(t *testing.T) {
	if c, _ := NewRequestGet(os.Getenv(`TEST_TARGET`)); c.GetMethod() != http.MethodGet {
		t.Error()
	}
}

func TestNewRequestHead(t *testing.T) {
	if c, _ := NewRequestHead(os.Getenv(`TEST_TARGET`)); c.GetMethod() != http.MethodHead {
		t.Error()
	}
}

func TestNewRequestPatch(t *testing.T) {
	if c, _ := NewRequestPatch(os.Getenv(`TEST_TARGET`)); c.GetMethod() != http.MethodPatch {
		t.Error()
	}
}

func TestNewRequestTrace(t *testing.T) {
	if c, _ := NewRequestTrace(os.Getenv(`TEST_TARGET`)); c.GetMethod() != http.MethodTrace {
		t.Error()
	}
}

func TestNewRequestPost(t *testing.T) {
	if c, _ := NewRequestPost(os.Getenv(`TEST_TARGET`)); c.GetMethod() != http.MethodPost {
		t.Error(c.GetMethod())
	}
}

func TestNewRequestPut(t *testing.T) {
	if c, _ := NewRequestPut(os.Getenv(`TEST_TARGET`)); c.GetMethod() != http.MethodPut {
		t.Error()
	}
}

func TestNewRequestConnect(t *testing.T) {
	if c, _ := NewRequestConnect(os.Getenv(`TEST_TARGET`)); c.GetMethod() != http.MethodConnect {
		t.Error()
	}
}

func TestNewRequestOptions(t *testing.T) {
	if c, _ := NewRequestOptions(os.Getenv(`TEST_TARGET`)); c.GetMethod() != http.MethodOptions {
		t.Error()
	}
}
