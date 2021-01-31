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
# requests
BenchmarkClient_GClientGet-12    	        2500	    814623 ns/op
BenchmarkClient_HttpClientGet-12    	    2024	    870806 ns/op

# new request
Benchmark_Gclient_NewRequest-12    	  799914	      1350 ns/op
Benchmark_Http_NewRequest-12    	  799903	      1321 ns/op
```