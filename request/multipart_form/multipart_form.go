package multipart_form

import (
	"io"
	"mime/multipart"
	"os"
)

type Option func(multipartWriter *multipart.Writer) (err error)

//File set file
func File(field, filePath string) Option {
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

//Field set field
func Field(field, value string) Option {
	return func(multipartWriter *multipart.Writer) (err error) {
		err = multipartWriter.WriteField(field, value)
		return
	}
}

//Boundary set boundary
func Boundary(boundary string) Option {
	return func(multipartWriter *multipart.Writer) (err error) {
		err = multipartWriter.SetBoundary(boundary)
		return
	}
}
