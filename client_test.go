package gclient

import (
	"github.com/jianzhiyao/gclient/tests"
	"net/http"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	if c := New(); c == nil {
		t.Error()
	}
}

func TestClient_Option(t *testing.T) {
	timeout := 5423 * time.Second
	c := New()
	c.Option(OptTimeout(timeout))

	if c.clientTimeout != timeout {
		t.Error()
	}
}

func TestClient_Options(t *testing.T) {
	timeout := 12398 * time.Second
	retry := 86543

	c := New()
	c.Options(
		OptTimeout(timeout),
		OptRetry(retry),
	)

	if c.clientTimeout != timeout {
		t.Error()
	}
	if c.retry != retry {
		t.Error()
	}
}

func TestClient_Do(t *testing.T) {
	c := New()
	url := tests.GetServerUrl()

	if resp, err := c.Do(http.MethodGet, url); err != nil {
		t.Error(err)
	} else if resp.StatusCode() != http.StatusOK {
		t.Error()
	}
}

func TestClient_DoRequest(t *testing.T) {
	c := New()
	url := tests.GetServerUrl()

	if req, err := NewRequest(http.MethodGet, url); err != nil {
		t.Error(err)
	} else {
		if resp, err := c.DoRequest(req); err != nil {
			t.Error(err)
		} else if resp.StatusCode() != http.StatusOK {
			t.Error(`resp.StatusCode()`, resp.StatusCode())
		}
	}

}

func BenchmarkClient_GClientGet(b *testing.B) {
	benchmarkLimit, _ := strconv.Atoi(os.Getenv(`BENCHMARK_LIMIT`))
	limit := make(chan bool, benchmarkLimit)

	c := New()
	url := os.Getenv(`BENCHMARK_TARGET`)
	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		limit <- true
		go func() {
			req, _ := NewRequestGet(url)
			if resp, err := c.DoRequest(req); err != nil {
				b.Error(err)
			} else {
				if resp == nil || resp.StatusCode() != http.StatusOK {
					b.Error()
				}
			}
			wg.Done()
			<-limit
		}()
	}
	wg.Wait()
}

func BenchmarkClient_HttpClientGet(b *testing.B) {
	benchmarkLimit, _ := strconv.Atoi(os.Getenv(`BENCHMARK_LIMIT`))
	limit := make(chan bool, benchmarkLimit)

	c := &http.Client{}
	url := os.Getenv(`BENCHMARK_TARGET`)
	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		limit <- true
		go func() {
			req, _ := http.NewRequest(http.MethodGet, url, nil)
			if resp, err := c.Do(req); err != nil {
				b.Error(err)
			} else {
				if resp == nil || resp.StatusCode != http.StatusOK {
					b.Error()
				}
			}
			wg.Done()
			<-limit
		}()
	}
	wg.Wait()
}

func Benchmark_Gclient_NewRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if req, err := NewRequest(http.MethodGet, tests.GetServerUrl()); err != nil {
			b.Error(err)
		} else if req == nil {
			b.Error()
		}
	}
}

func Benchmark_Http_NewRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if req, err := http.NewRequest(http.MethodGet, tests.GetServerUrl(), nil); err != nil {
			b.Error(err)
		} else if req == nil {
			b.Error()
		}
	}
}
