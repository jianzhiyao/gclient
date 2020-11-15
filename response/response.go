package response

import (
	"compress/flate"
	"compress/gzip"
	"errors"
	"github.com/dsnet/compress/brotli"
	"io/ioutil"
	"net/http"
)

type Response struct {
	resp *http.Response
}

func New(response *http.Response) *Response {
	return &Response{resp: response}
}

func (r *Response) readBytes() ([]byte, error) {
	var body []byte
	if r.resp == nil || r.resp.Body == nil {
		return nil, errors.New(`invalid response`)
	}
	defer r.resp.Body.Close()

	switch r.resp.Header.Get("Content-Encoding") {
	case `gzip`:
		reader, err := gzip.NewReader(r.resp.Body)
		defer reader.Close()
		if err != nil {
			return nil, err
		}
		body, err = ioutil.ReadAll(reader)
		return body, err
	case `deflate`:
		reader := flate.NewReader(r.resp.Body)
		defer reader.Close()
		body, err := ioutil.ReadAll(reader)
		return body, err
	case `br`:
		reader, err := brotli.NewReader(r.resp.Body, nil)
		if err != nil {
			return nil, err
		}
		defer reader.Close()
		body, err := ioutil.ReadAll(reader)
		return body, err
	}

	body, err := ioutil.ReadAll(r.resp.Body)
	return body, err
}

func (r *Response) JsonUnmarshal(v interface{}) error {
	return nil
}

func (r *Response) XmlUnmarshal(v interface{}) error {
	return nil
}

func (r *Response) YamlUnmarshal(v interface{}) error {
	return nil
}
