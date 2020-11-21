package mutipart_form

import (
	"bytes"
	"github.com/google/uuid"
	"mime/multipart"
	"strings"
	"testing"
)

func TestFile(t *testing.T) {
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)

	filePath1 := `test_files/test_file1.txt`
	filePath2 := `test_files/test_file2.txt`
	options := []Option{
		File("file1", filePath1),
		File("file2", filePath2),
	}

	for _, opt := range options {
		_ = opt(bodyWriter)
	}

	if !strings.Contains(bodyBuffer.String(), filePath1) {
		t.Error()
		return
	}

	if !strings.Contains(bodyBuffer.String(), filePath2) {
		t.Error()
		return
	}
}

func TestField(t *testing.T) {
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)

	uuid1 := uuid.New().String()
	uuid2 := uuid.New().String()
	options := []Option{
		Field("id1", uuid1),
		Field("id2", uuid2),
	}

	for _, opt := range options {
		_ = opt(bodyWriter)
	}

	if !strings.Contains(bodyBuffer.String(), uuid1) {
		t.Error()
		return
	}

	if !strings.Contains(bodyBuffer.String(), uuid2) {
		t.Error()
		return
	}

}

func TestBoundary(t *testing.T) {
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)

	boundary1 := uuid.New().String()
	filePath1 := `test_files/test_file1.txt`
	filePath2 := `test_files/test_file2.txt`
	options1 := []Option{
		Boundary(boundary1),
		File("file1", filePath1),
		File("file2", filePath2),
		Field("aaaa", `1`),
	}

	for _, opt := range options1 {
		_ = opt(bodyWriter)
	}

	if c := strings.Count(bodyBuffer.String(), boundary1); c != 3 {
		t.Error(c)
		return
	}

	boundary2 := uuid.New().String()
	options2 := []Option{
		Boundary(boundary2),
	}

	for _, opt := range options2 {
		_ = opt(bodyWriter)
	}

	if c := strings.Count(bodyBuffer.String(), boundary1); c != 3 {
		t.Error(c)
		return
	}
}
