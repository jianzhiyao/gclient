package client

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

type Client struct {
	ctx context.Context

	retry int
	//Client level headers
	headers map[string]string

	clientCookieJar     *cookiejar.Jar
	clientTransport     http.RoundTripper
	clientCheckRedirect CheckRedirectHandler
	clientTimeout       time.Duration

	sign int8
}

func (r *Client) Option(option Option) *Client {
	option(r)
	return r
}

func (r *Client) Close() {

}

func (r *Client) Options(options ...Option) *Client {
	for _, option := range options {
		r.Option(option)
	}
	return r
}

func (r *Client) newHttpClient() (c *http.Client, returnBack ReturnHttpClient) {
	c, returnBack = getClientFromPool()
	c.Transport = r.clientTransport
	c.CheckRedirect = r.clientCheckRedirect
	c.Jar = r.clientCookieJar
	c.Timeout = r.clientTimeout

	return
}

func (r *Client) Do(method, url string, body io.Reader) (*response.Response, error) {
	var (
		resp *http.Response
		err  error
	)
	c, returnBack := r.newHttpClient()
	defer returnBack(c)

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
		r.headers[structs.HeaderAcceptEncoding] = strings.Join(contentEncoding, `, `)
	}

	for key, value := range r.headers {
		req.Header.Set(key, value)
	}

	tryCount := r.retry
	if tryCount <= 1 {
		tryCount = 1
	}
	for tryCount > 0 {
		resp, err = c.Do(req)
		if err != nil {
			break
		}
		tryCount--
	}
	if err != nil {
		return nil, err
	}

	return response.New(resp), nil
}

func New(options ...Option) *Client {
	req := &Client{
		ctx:                 nil,
		retry:               0,
		headers:             make(map[string]string),
		clientCookieJar:     nil,
		clientTransport:     nil,
		clientCheckRedirect: nil,
		clientTimeout:       0,
		sign:                0,
	}

	req.Options(
		OptCookieJar(nil),
		OptTransport(nil),
		OptCheckRedirectHandler(nil),
	)

	return req.Options(options...)
}
