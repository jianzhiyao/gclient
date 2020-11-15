package request

import (
	"context"
	"github.com/jianzhiyao/gclient/response"
	"github.com/jianzhiyao/gclient/structs"
	"io"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"
)

type Sign int8

const (
	SignGzip    Sign = 1 << 0
	SignDeflate Sign = 1 << 1
	SignBr      Sign = 1 << 2
)

type Request struct {
	ctx context.Context

	retry int
	//Request headers
	headers map[string]string

	clientCookieJar     *cookiejar.Jar
	clientTransport     http.RoundTripper
	clientCheckRedirect CheckRedirectHandler
	clientTimeout       time.Duration

	sign int8
}

func (r *Request) Option(option Option) *Request {
	option(r)
	return r
}

func (r *Request) Close() {

}

func (r *Request) Options(options ...Option) *Request {
	for _, option := range options {
		r.Option(option)
	}
	return r
}

func (r *Request) newHttpClient() (c *http.Client, putBack func(client *http.Client)) {
	c = getClientFromPool()
	c.Transport = r.clientTransport
	c.CheckRedirect = r.clientCheckRedirect
	c.Jar = r.clientCookieJar
	c.Timeout = r.clientTimeout



	return c, putClientToPool
}

func (r *Request) Do(method, url string, body io.Reader) (*response.Response, error) {
	c, putBack := r.newHttpClient()
	defer putBack(c)
	//
	//c := &http.Client{
	//	Transport:     nil,
	//	CheckRedirect: nil,
	//	Jar:           nil,
	//	Timeout:       0,
	//}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	var contentEncoding []string
	if r.sign&int8(SignGzip) != 0 {
		contentEncoding = append(contentEncoding, structs.ContentEncodingGzip)
	}
	if r.sign&int8(SignDeflate) != 0 {
		contentEncoding = append(contentEncoding, structs.ContentEncodingDeflate)
	}
	if r.sign&int8(SignBr) != 0 {
		contentEncoding = append(contentEncoding, structs.ContentEncodingBr)
	}
	if len(contentEncoding) > 0 {
		r.headers[HeaderAcceptEncoding] = strings.Join(contentEncoding, `, `)
	}

	for key, value := range r.headers {
		req.Header.Set(key, value)
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	return response.New(resp), nil
}

func New(options ...Option) *Request {
	req := &Request{
		ctx:                 nil,
		retry:               0,
		clientTimeout:       0,
		headers:             make(map[string]string),
		clientCookieJar:     nil,
		clientTransport:     nil,
		clientCheckRedirect: nil,
		sign:                0,
	}
	return req.Options(options...)
}
