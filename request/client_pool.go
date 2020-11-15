package request

import (
	"net/http"
	"sync"
)

var httpClientPool *sync.Pool
var httpClientOnce sync.Once

func init() {
	httpClientOnce.Do(func() {
		httpClientPool = &sync.Pool{
			New: func() interface{} {
				return &http.Client{}
			},
		}
	})
}

func getClientFromPool() *http.Client {
	cli := httpClientPool.Get().(*http.Client)

	return cli
}

func putClientToPool(cli *http.Client) {
	cli.Timeout = 0
	cli.Transport = nil
	cli.CheckRedirect = nil
	cli.Jar = nil

	httpClientPool.Put(cli)
}
