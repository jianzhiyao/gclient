package request

import (
	"context"
	"net/http"
	"testing"
	"time"
)

func TestOptContext(t *testing.T) {
	ctx := context.WithValue(context.Background(), `1`, `2`)
	req := New(OptContext(ctx))

	if req.ctx.Value(`1`) != `2` {
		t.Error()
		return
	}
}

func TestOptEnableBr(t *testing.T) {
	req := New(OptEnableBr())

	if req.sign&int8(SignBr) == 0 {
		t.Error()
		return
	}
}

func TestOptDisableBr(t *testing.T) {
	req := New(
		OptEnableBr(),
		OptDisableBr(),
	)

	if req.sign&int8(SignBr) != 0 {
		t.Error()
		return
	}
}

func TestOptEnableGzip(t *testing.T) {
	req := New(OptEnableGzip())

	if req.sign&int8(SignGzip) == 0 {
		t.Error()
		return
	}
}

func TestOptDisableGzip(t *testing.T) {
	req := New(
		OptEnableGzip(),
		OptDisableGzip(),
	)

	if req.sign&int8(SignGzip) != 0 {
		t.Error()
		return
	}
}

func TestOptEnableDeflate(t *testing.T) {
	req := New(OptEnableDeflate())

	if req.sign&int8(SignDeflate) == 0 {
		t.Error()
		return
	}
}

func TestOptDisableDeflate(t *testing.T) {
	req := New(
		OptEnableDeflate(),
		OptDisableDeflate(),
	)

	if req.sign&int8(SignDeflate) != 0 {
		t.Error()
		return
	}
}

func TestOptTransport(t *testing.T) {
	req := New(
		OptTransport(&http.Transport{}),
	)

	if req.clientTransport == nil {
		t.Error()
		return
	}
}

func TestOptCookieJar(t *testing.T) {
	req := New(
		OptCookieJar(nil),
	)

	if req.clientCookieJar == nil {
		t.Error()
		return
	}
}

func TestOptEnableCheckRedirect(t *testing.T) {
	req := New(
		OptEnableCheckRedirect(func(req *http.Request, via []*http.Request) error {
			return nil
		}),
	)

	if req.clientCheckRedirect == nil {
		t.Error()
		return
	}
}

func TestOptHeader(t *testing.T) {
	req := New(
		OptHeader(`a`, `b`),
		OptHeader(`c`, `d`),
		OptHeader(`e`, `f`),
	)

	if req.headers[`a`] != `b` {
		t.Error()
		return
	}
	if req.headers[`c`] != `d` {
		t.Error()
		return
	}
	if req.headers[`e`] != `f` {
		t.Error()
		return
	}
}

func TestOptHeaders(t *testing.T) {
	req := New(
		OptHeaders(map[string]string{
			`a`: `b`,
			`c`: `d`,
			`e`: `f`,
		}),
	)

	if req.headers[`a`] != `b` {
		t.Error()
		return
	}
	if req.headers[`c`] != `d` {
		t.Error()
		return
	}
	if req.headers[`e`] != `f` {
		t.Error()
		return
	}
}

func TestOptRetry(t *testing.T) {
	req := New(
		OptRetry(88554),
	)

	if req.retry != 88554 {
		t.Error()
		return
	}
}

func TestOptUserAgent(t *testing.T) {
	req := New(
		OptUserAgent(`server`),
	)

	if req.headers[HeaderUserAgent] != `server` {
		t.Error()
		return
	}
}

func TestOptTimeout(t *testing.T) {
	req := New(
		OptTimeout(878 * time.Second),
	)

	if req.clientTimeout != 878*time.Second {
		t.Error()
		return
	}
}