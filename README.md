# gclient
[![Build Status](https://travis-ci.com/jianzhiyao/gclient.svg?branch=master)](https://travis-ci.com/jianzhiyao/gclient) [![GoDoc](http://godoc.org/github.com/jianzhiyao/gclient?status.svg)](http://godoc.org/github.com/jianzhiyao/gclient) [![Foundation](https://img.shields.io/badge/Golang-Foundation-green.svg)](http://golangfoundation.org) [![Go Report Card](https://goreportcard.com/badge/github.com/jianzhiyao/gclient)](https://goreportcard.com/report/github.com/jianzhiyao/gclient)

http.Client & http.Request implementation for golang

# documentation
[go to ducumentation](https://pkg.go.dev/github.com/jianzhiyao/gclient)

## How to use gclient

``go get github.com/jianzhiyao/gclient``

## Client
### New a client

```go
cli := gclient.New(
		//with context
		gclient.OptContext(context.Background()),
		//set timeout
		gclient.OptTimeout(3*time.Second),
		//retry if get error after requesting
		gclient.OptRetry(3),
		//enable compression: br
		gclient.OptEnableBr(),
		//enable compression: gzip
		gclient.OptEnableGzip(),
		//enable compression: deflate
		gclient.OptEnableDeflate(),
		//set header for client level
		gclient.OptHeader("token", "request_token"),
		gclient.OptHeaders(map[string][]string{
			`accept-language`: []string{
				`zh-CN`,
				`zh;q=0.9`,
				`en;q=0.8`,
				`en-US;q=0.7`,
			},
		}),
		//set cookie jar for http.Client
		gclient.OptCookieJar(nil),
		//set transport for http.Client
		gclient.OptTransport(nil),
        //resize worker poll size(default 1000)
		gclient.OptWorkerPoolSize(5),
	)
```

## Request
### Simple request
```go
cli := gclient.New()
if resp, err := cli.Do(http.MethodHead, "http://exmaple.com/job.json"); err != nil {
	fmt.Println(err)
} else {
	fmt.Println(resp.String())
}
```

### Complex request

supported methods
- NewRequest
- NewRequestGet
- NewRequestHead
- NewRequestPost
- NewRequestPut
- NewRequestPatch
- NewRequestDelete
- NewRequestConnect
- NewRequestOptions
- NewRequestTrace

```go
if req, err := gclient.NewRequest(http.MethodGet, "https://exmaple.com"); err != nil {
	fmt.Println(err)
	} else {
	var data Data
	if err := req.Json(&Data{}); err != nil {
		fmt.Println(err)
	} else {
		resp, err := cli.DoRequest(req)
	}
}
```
## Response

### Status
```go
fmt.Println(resp.Status())
fmt.Println(resp.StatusCode())
```

### Body
#### Get response content as string
```go
if body, err := resp.String(); err != nil {
	fmt.Println(err)
} else {
	fmt.Println(body)
}
```

#### Get response content as []byte
```go
if body, err := resp.Bytes(); err != nil {
	fmt.Println(err)
} else {
	fmt.Println(body)
}
```

#### Unmarshal response content as json
```go
var a Resp
if err := resp.JsonUnmarshal(&a); err != nil {
	fmt.Println(err)
} else {
	fmt.Println(a)
}
```

#### Unmarshal response content as yaml
```go
var a Resp
if err := resp.YamlUnmarshal(&a); err != nil {
	fmt.Println(err)
} else {
	fmt.Println(a)
}
```

#### Unmarshal response content as xml
```go
var a Resp
if err := resp.XmlUnmarshal(&a); err != nil {
	fmt.Println(err)
} else {
	fmt.Println(a)
}
```

## Benchmark
Gclient VS. net/http.Client
```
# test-1
BenchmarkClient_GClientGet-2      	     373	  14990893 ns/op	 5066127 B/op	   42150 allocs/op
BenchmarkClient_HttpClientGet-2   	     361	   6489557 ns/op	 2250161 B/op	   18406 allocs/op

# test-2
BenchmarkClient_GClientGet-2      	     267	  14451088 ns/op	 4812062 B/op	   40183 allocs/op
BenchmarkClient_HttpClientGet-2   	     320	   9057752 ns/op	 3043013 B/op	   25254 allocs/op

# test-3
BenchmarkClient_GClientGet-2      	     184	  11190623 ns/op	 3908556 B/op	   32397 allocs/op
BenchmarkClient_HttpClientGet-2   	     307	  12442205 ns/op	 4197268 B/op	   34799 allocs/op

# test-4
BenchmarkClient_GClientGet-2      	     231	   7910001 ns/op	 2671297 B/op	   22118 allocs/op
BenchmarkClient_HttpClientGet-2   	     334	  14316617 ns/op	 4900196 B/op	   40913 allocs/op
```