package multipart_form

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"mime/multipart"
	"os"
	"strings"
	"testing"
)

func TestFile(t *testing.T) {
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)

	filePath1 := fmt.Sprintf(`%s.txt`, uuid.New().String())
	filePath2 := fmt.Sprintf(`%s.txt`, uuid.New().String())

	file1, _ := os.Create(filePath1)
	file2, _ := os.Create(filePath2)

	content1 := uuid.New().String()
	content2 := uuid.New().String()
	_, _ = fmt.Fprintln(file1, content1)
	_, _ = fmt.Fprintln(file2, content2)

	defer func() {
		file1.Close()
		os.Remove(file1.Name())

		file2.Close()
		os.Remove(file2.Name())
	}()

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

	if !strings.Contains(bodyBuffer.String(), content1) {
		t.Error()
		return
	}

	if !strings.Contains(bodyBuffer.String(), content2) {
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

	filePath1 := fmt.Sprintf(`%s.txt`, uuid.New().String())
	filePath2 := fmt.Sprintf(`%s.txt`, uuid.New().String())

	file1, _ := os.Create(filePath1)
	file2, _ := os.Create(filePath2)

	content1 := uuid.New().String()
	content2 := uuid.New().String()
	_, _ = fmt.Fprintln(file1, content1)
	_, _ = fmt.Fprintln(file2, content2)
	defer func() {
		_ = file1.Close()
		_ = os.Remove(file1.Name())

		_ = file2.Close()
		_ = os.Remove(file2.Name())
	}()

	boundary1 := uuid.New().String()
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
