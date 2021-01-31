package request

import (
	"bytes"
	"encoding"
	"github.com/jianzhiyao/gclient/consts"
	"github.com/jianzhiyao/gclient/consts/content_type"
	"github.com/jianzhiyao/gclient/consts/transfer_encoding"
	"github.com/jianzhiyao/gclient/request/form"
	"github.com/jianzhiyao/gclient/request/multipart_form"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
)

type Request struct {
	request *http.Request
}

//New base method of new request
func New(method, uri string) (*Request, error) {

	if req, err := http.NewRequest(method, uri, nil); err != nil {
		return nil, err
	} else {
		return &Request{
			request: req,
		}, nil
	}
}

func (r *Request) SetHeader(key string, value ...string) {
	if r.request.Header == nil {
		r.request.Header = http.Header{}
	}
	r.request.Header[key] = value
}

func (r *Request) GetUrl() string {
	return r.request.URL.String()
}

func (r *Request) GetMethod() string {
	return r.request.Method
}

func (r *Request) GetHeaders() http.Header {
	return r.request.Header
}

func (r *Request) GetHeader(key string) (value []string, ok bool) {
	if r.request.Header == nil {
		return
	}

	value, ok = r.request.Header[key]
	return
}

func (r *Request) GetBody() io.Reader {
	return r.request.Body
}

func (r *Request) Json(body interface{}) (err error) {
	if e := r.Body(body); e != nil {
		return e
	}

	r.SetHeader(consts.HeaderContentType, content_type.ApplicationJson)
	return
}

func (r *Request) Xml(body interface{}) (err error) {
	if e := r.Body(body); e != nil {
		return e
	}
	r.SetHeader(consts.HeaderContentType, content_type.ApplicationXml)
	return
}

func (r *Request) MultiForm(options ...multipart_form.Option) (err error) {
	pr, pw := io.Pipe()

	bodyWriter := multipart.NewWriter(pw)

	go func() {
		defer func() {
			bodyWriter.Close()
			pw.Close()
		}()
		for _, option := range options {
			_ = option(bodyWriter)
		}
	}()

	if e := r.Body(pr); e != nil {
		return
	}

	r.SetHeader(consts.HeaderContentType, bodyWriter.FormDataContentType())
	r.SetHeader(consts.HeaderTransferEncoding, transfer_encoding.Chunked)
	return
}

func (r *Request) Form(options ...form.Option) (err error) {
	values := url.Values{}
	for _, option := range options {
		if e := option(values); e != nil {
			return e
		}
	}
	if e := r.Body(values.Encode()); e != nil {
		return
	}

	r.SetHeader(consts.HeaderContentType, content_type.ApplicationXWwwFormUrlencoded)
	return
}

func (r *Request) Body(body interface{}) (err error) {
	var reader io.Reader
	switch body := body.(type) {
	case []byte:
		reader = bytes.NewReader(body)
	case string:
		reader = bytes.NewReader([]byte(body))
	case encoding.BinaryMarshaler:
		if bodyBytes, e := body.MarshalBinary(); e != nil {
			err = e
		} else {
			reader = bytes.NewReader(bodyBytes)
		}
	case io.Reader:
		reader = body
	default:
		err = ErrCanNotMarshal
	}

	if err != nil {
		return
	}
	r.request.Body = ioutil.NopCloser(reader)

	return
}
