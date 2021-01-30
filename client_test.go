package gclient

import (
	"net/http"
	"os"
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
	url := os.Getenv(`TEST_TARGET`)

	if resp, err := c.Do(http.MethodGet, url); err != nil {
		t.Error(err)
	} else if resp.StatusCode() != http.StatusOK {
		t.Error()
	}
}

func TestClient_DoRequest(t *testing.T) {
	c := New()
	url := os.Getenv(`TEST_TARGET`)

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

const BenchmarkLimit = 1000
func benchmarkWithWorker(b *testing.B, size int) {
	limit := make(chan bool, BenchmarkLimit)

	c := New(
		OptWorkerPoolSize(size),
	)
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

func BenchmarkClient_GClientGet_1_Workers(b *testing.B) {
	benchmarkWithWorker(b, 1)
}

func BenchmarkClient_GClientGet_10_Workers(b *testing.B) {
	benchmarkWithWorker(b, 10)
}

func BenchmarkClient_GClientGet_100_Workers(b *testing.B) {
	benchmarkWithWorker(b, 100)
}

func BenchmarkClient_GClientGet_1000_Workers(b *testing.B) {
	benchmarkWithWorker(b, 1000)
}

func BenchmarkClient_HttpClientGet(b *testing.B) {
	limit := make(chan bool, BenchmarkLimit)

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
