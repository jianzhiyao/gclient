package request

import (
	"context"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type Option func(req *Request)
type CheckRedirectHandler func(req *http.Request, via []*http.Request) error

func OptTimeout(timeout time.Duration) Option {
	return func(req *Request) {
		req.clientTimeout = timeout
	}
}

func OptContext(ctx context.Context) Option {
	return func(req *Request) {
		req.ctx = ctx
	}
}

func OptHeader(key, value string) Option {
	return func(req *Request) {
		req.headers[key] = value
	}
}

func OptUserAgent(ua string) Option {
	return OptHeader(HeaderUserAgent, ua)
}

func OptHeaders(headers map[string]string) Option {
	return func(req *Request) {
		for key, value := range headers {
			req.headers[key] = value
		}
	}
}

func OptEnableGzip() Option {
	return enableSign(SignGzip)
}

func OptDisableGzip() Option {
	return disableSign(SignGzip)
}

func OptEnableDeflate() Option {
	return enableSign(SignDeflate)
}

func OptDisableDeflate() Option {
	return disableSign(SignDeflate)
}

func OptEnableBr() Option {
	return enableSign(SignBr)
}

func OptDisableBr() Option {
	return disableSign(SignBr)
}

func OptEnableCookie(options *cookiejar.Options) Option {
	return func(req *Request) {
		var err error
		req.clientCookieJar, err = cookiejar.New(options)
		if err != nil {
			panic(err)
		}
	}
}

func OptDisableCookie() Option {
	return func(req *Request) {
		req.clientCookieJar = nil
	}
}

func OptEnableTransport(roundTripper http.RoundTripper) Option {
	return func(req *Request) {
		req.clientTransport = roundTripper
	}
}

func OptDisableTransport() Option {
	return func(req *Request) {
		req.clientTransport = nil
	}
}

func OptEnableCheckRedirect(clientCheckRedirect CheckRedirectHandler) Option {
	return func(req *Request) {
		req.clientCheckRedirect = clientCheckRedirect
	}
}

func OptDisableCheckRedirect() Option {
	return func(req *Request) {
		req.clientCheckRedirect = nil
	}
}

func OptRetry(times int) Option {
	return func(req *Request) {
		req.retry = times
	}
}

func enableSign(t Sign) Option {
	return func(req *Request) {
		req.sign |= int8(t)
	}
}

func disableSign(t Sign) Option {
	return func(req *Request) {
		req.sign ^= int8(t)
	}
}
