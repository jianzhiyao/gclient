package gclient

import (
	"context"
	"github.com/jianzhiyao/gclient/consts"
	"github.com/jianzhiyao/gclient/request"
	"github.com/jianzhiyao/gclient/response"
	"io"
	"net/http"
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
	headers http.Header

	clientCookieJar     http.CookieJar
	clientTransport     http.RoundTripper
	clientCheckRedirect CheckRedirectHandler
	clientTimeout       time.Duration

	sign int8
}

func New(options ...Option) *Client {
	c := &Client{
		headers: http.Header{},
	}

	c.Options(options...)

	return c
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

func (r *Client) Do(method, url string) (*response.Response, error) {
	return r.do(method, url, nil, nil)
}

func (r *Client) DoRequest(req *request.Request) (*response.Response, error) {
	return r.do(
		req.GetMethod(),
		req.GetUrl(),
		req.GetBody(),
		req.GetHeaders(),
	)
}

func (r *Client) do(method, url string, body io.Reader, headers http.Header) (*response.Response, error) {
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
		contentEncoding = append(contentEncoding, consts.ContentEncodingGzip)
	}
	if r.sign&int8(SignDeflate) != 0 {
		contentEncoding = append(contentEncoding, consts.ContentEncodingDeflate)
	}
	if r.sign&int8(SignBr) != 0 {
		contentEncoding = append(contentEncoding, consts.ContentEncodingBr)
	}
	if len(contentEncoding) > 0 {
		r.headers[consts.HeaderAcceptEncoding] = contentEncoding
	}

	//set request headers
	//header from client
	req.Header = r.headers
	//header from request
	for key, header := range headers {
		req.Header[key] = header
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

	return response.New(resp)
}
