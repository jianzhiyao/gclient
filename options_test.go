package gclient

import (
	"context"
	"github.com/jianzhiyao/gclient/consts"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
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

func TestOptEnableBr2(t *testing.T) {
	req := New(
		OptEnableBr(),
	)
	url := os.Getenv(`TEST_TARGET`)

	resp, err := req.Do(http.MethodGet, url)
	if err != nil {
		t.Error(err)
		return
	}

	if !strings.Contains(resp.Header(consts.HeaderContentEncoding), consts.ContentEncodingBr) {
		t.Error()
		return
	}

	_, err = resp.String()
	if err != nil {
		t.Error(err)
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

func TestOptEnableGzip2(t *testing.T) {
	req := New(
		OptEnableGzip(),
	)
	url := os.Getenv(`TEST_TARGET`)

	resp, err := req.Do(http.MethodGet, url)
	if err != nil {
		t.Error(err)
		return
	}

	if !strings.Contains(resp.Header(consts.HeaderContentEncoding), consts.ContentEncodingGzip) {
		t.Error()
		return
	}
	_, err = resp.String()
	if err != nil {
		t.Error(err)
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

func TestOptEnableDeflate2(t *testing.T) {
	req := New(
		OptEnableDeflate(),
	)
	url := os.Getenv(`TEST_TARGET`)

	resp, err := req.Do(http.MethodGet, url)
	if err != nil {
		t.Error(err)
		return
	}

	if !strings.Contains(resp.Header(consts.HeaderContentEncoding), consts.ContentEncodingDeflate) {
		t.Error()
		return
	}
	_, err = resp.String()
	if err != nil {
		t.Error(err)
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

	if req.clientCookieJar != nil {
		t.Error()
		return
	}
}

func TestOptCookieJar1(t *testing.T) {
	if jar, err := cookiejar.New(nil); err != nil {
		t.Error()
		return
	} else {
		req := New(
			OptCookieJar(jar),
		)

		if req.clientCookieJar == nil {

		}
	}
}

func TestOptCheckRedirectHandler(t *testing.T) {
	req := New(
		OptCheckRedirectHandler(func(req *http.Request, via []*http.Request) error {
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

	if req.headers[`a`][0] != `b` {
		t.Error()
		return
	}
	if req.headers[`c`][0] != `d` {
		t.Error()
		return
	}
	if req.headers[`e`][0] != `f` {
		t.Error()
		return
	}
}

func TestOptHeaders(t *testing.T) {
	req := New(
		OptHeaders(map[string][]string{
			`a`: {`b`},
			`c`: {`d`},
			`e`: {`f`},
		}),
	)

	if req.headers[`a`][0] != `b` {
		t.Error()
		return
	}
	if req.headers[`c`][0] != `d` {
		t.Error()
		return
	}
	if req.headers[`e`][0] != `f` {
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

	if req.headers[consts.HeaderUserAgent][0] != `server` {
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

func TestOptWorkerPoolSize(t *testing.T) {
	req := New(
		OptWorkerPoolSize(12543),
	)

	if req.pool.Cap() != 12543 {
		t.Error()
		return
	}
}