package request

import (
	"context"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type Option func(req *Request)
type CheckRedirectHandler func(req *http.Request, via []*http.Request) error

func Timeout(timeout time.Duration) Option {
	return func(req *Request) {
		req.clientTimeout = timeout
	}
}

func Context(ctx context.Context) Option {
	return func(req *Request) {
		req.ctx = ctx
	}
}

func Header(key, value string) Option {
	return func(req *Request) {
		req.headers[key] = value
	}
}

func UserAgent(ua string) Option {
	return Header(HeaderUserAgent, ua)
}

func Headers(headers map[string]string) Option {
	return func(req *Request) {
		for key, value := range headers {
			req.headers[key] = value
		}
	}
}

func EnableGzip() Option {
	return enableSign(SignGzip)
}

func DisableGzip() Option {
	return disableSign(SignGzip)
}

func EnableDeflate() Option {
	return enableSign(SignDeflate)
}

func DisableDeflate() Option {
	return disableSign(SignDeflate)
}

func EnableBr() Option {
	return enableSign(SignBr)
}

func DisableBr() Option {
	return disableSign(SignBr)
}

func EnableCookie(options *cookiejar.Options) Option {
	return func(req *Request) {
		var err error
		req.clientCookieJar, err = cookiejar.New(options)
		if err != nil {
			panic(err)
		}
	}
}

func DisableCookie() Option {
	return func(req *Request) {
		req.clientCookieJar = nil
	}
}

func EnableTransport(roundTripper http.RoundTripper) Option {
	return func(req *Request) {
		req.clientTransport = roundTripper
	}
}

func DisableTransport() Option {
	return func(req *Request) {
		req.clientTransport = nil
	}
}

func EnableCheckRedirect(clientCheckRedirect CheckRedirectHandler) Option {
	return func(req *Request) {
		req.clientCheckRedirect = clientCheckRedirect
	}
}

func DisableCheckRedirect() Option {
	return func(req *Request) {
		req.clientCheckRedirect = nil
	}
}

func Retry(times int) Option {
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
