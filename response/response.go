package response

import (
	"compress/flate"
	"compress/gzip"
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/dsnet/compress/brotli"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"strings"
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

	contentEncoding := r.resp.Header.Get("Content-Encoding")
	if strings.Contains(contentEncoding, `gzip`) {
		reader, err := gzip.NewReader(r.resp.Body)
		defer reader.Close()
		if err != nil {
			return nil, err
		}
		body, err = ioutil.ReadAll(reader)
		return body, err
	} else if strings.Contains(contentEncoding, `deflate`) {
		reader := flate.NewReader(r.resp.Body)
		defer reader.Close()
		body, err := ioutil.ReadAll(reader)
		return body, err
	} else if strings.Contains(contentEncoding, `br`) {
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
	b, err := r.readBytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}

func (r *Response) XmlUnmarshal(v interface{}) error {
	b, err := r.readBytes()
	if err != nil {
		return err
	}
	return xml.Unmarshal(b, v)
}

func (r *Response) YamlUnmarshal(v interface{}) error {
	b, err := r.readBytes()
	if err != nil {
		return err
	}
	return yaml.Unmarshal(b, v)
}

func (r *Response) String() (string, error) {
	b, err := r.readBytes()
	if err != nil {
		return ``, err
	}
	return string(b), nil
}

func (r *Response) Header(key string) string {
	return r.resp.Header.Get(key)
}
func (r *Response) Headers() map[string][]string {
	m := make(map[string][]string)
	for key, value := range r.resp.Header {
		m[key] = value
	}
	return m
}
