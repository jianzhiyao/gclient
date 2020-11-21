package request

import (
	"io"
	"mime/multipart"
	"os"
)

type FormOption func(multipartWriter *multipart.Writer) (err error)

func File(field, filePath string) FormOption {
	return func(multipartWriter *multipart.Writer) (err error) {

		fileWriter, err := multipartWriter.CreateFormFile(field, filePath)
		if err != nil {
			return
		}
		file, err := os.Open(filePath)
		if err != nil {
			return
		}
		defer file.Close()

		if _, e := io.Copy(fileWriter, file); e != nil {
			return e
		}
		return
	}
}

func Field(field, value string) FormOption {
	return func(multipartWriter *multipart.Writer) (err error) {
		err = multipartWriter.WriteField(field, value)
		return
	}
}

func Boundary(boundary string) FormOption {
	return func(multipartWriter *multipart.Writer) (err error) {
		err = multipartWriter.SetBoundary(boundary)
		return
	}
}