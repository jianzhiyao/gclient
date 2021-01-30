package response

import (
	"compress/gzip"
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/dsnet/compress/brotli"
	"github.com/jianzhiyao/gclient/consts"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"net/http"
)

type Response struct {
	resp *http.Response
}

func New(response *http.Response) (*Response, error) {
	if response == nil {
		return nil, errors.New("not a valid response")

	}
	return &Response{resp: response}, nil
}

func (r *Response) Status() string {
	return r.resp.Status
}
func (r *Response) StatusCode() int {
	return r.resp.StatusCode
}

func (r *Response) readBytes() ([]byte, error) {
	if r.resp.Body == nil {
		return nil, errors.New("invalid response body")
	}
	defer r.resp.Body.Close()

	var (
		reader io.ReadCloser
		err    error
	)
	switch r.resp.Header.Get(consts.HeaderContentEncoding) {
	case consts.ContentEncodingGzip:
		reader, err = gzip.NewReader(r.resp.Body)
	case consts.ContentEncodingBr:
		reader, err = brotli.NewReader(r.resp.Body, nil)
	default:
		reader = r.resp.Body
	}
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(reader)
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

func (r *Response) Bytes() ([]byte, error) {
	return r.readBytes()
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
