package tests

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mholt/archiver/v3"
	"net/http"
	"net/http/httptest"
	"strings"
)

var server = newServer()

func GetServer() *httptest.Server {
	return server
}

func GetServerUrl() string {
	return GetServer().URL + "/"
}

func newServer() *httptest.Server {
	router := gin.Default()
	api := router.Use(
		gin.Recovery(),
	)
	{
		//UT
		ut := api.Use(
			compression(),
		)
		{
			ut.Any(`/`, ok())
			ut.GET(`/ok`, ok())
			ut.GET(`/json`, func(context *gin.Context) {
				context.JSON(http.StatusOK, gin.H{
					"code":    1,
					"message": "ok",
					"users":   []string{`aaron`, `john`},
				})
			})
			ut.GET(`/xml`, func(context *gin.Context) {
				context.XML(http.StatusOK, gin.H{
					"message": "ok",
				})
			})
			ut.GET(`/yaml`, func(context *gin.Context) {
				context.YAML(http.StatusOK, gin.H{
					"code":    1,
					"message": "ok",
					"users":   []string{`aaron`, `john`},
				})
			})
		}

		//benchmark
		api.Any(`/benchmark`, ok())
	}

	return httptest.NewServer(router)
}

func compression() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var compressor archiver.Compressor
		acceptEncoding := ctx.GetHeader(`Accept-Encoding`)
		encoding := ``
		if strings.Contains(acceptEncoding, `br`) {
			encoding = `br`
			compressor = archiver.NewBrotli()
		} else if strings.Contains(acceptEncoding, `gzip`) {
			encoding = `gzip`
			compressor = archiver.NewGz()
		}

		if compressor != nil {
			ctx.Writer = &handler{
				ResponseWriter: ctx.Writer,
				Compressor:     compressor,
			}
			ctx.Header("Content-Encoding", encoding)
			defer func() {
				if encoding != `` {
					ctx.Header("Content-Length", fmt.Sprint(ctx.Writer.Size()))
				}
			}()
		}

		ctx.Next()
	}
}

type handler struct {
	gin.ResponseWriter
	archiver.Compressor
}

func (g *handler) WriteString(s string) (count int, err error) {
	g.Header().Del("Content-Length")

	if g.Compressor != nil {
		count = len(s)
		reader := bytes.NewBufferString(s)
		err = g.Compressor.Compress(reader, g.ResponseWriter)
	} else {
		return g.ResponseWriter.WriteString(s)
	}

	return
}

func (g *handler) Write(data []byte) (count int, err error) {
	g.Header().Del("Content-Length")

	if g.Compressor != nil {
		count = len(data)
		reader := bytes.NewBuffer(data)
		err = g.Compressor.Compress(reader, g.ResponseWriter)
	} else {
		return g.ResponseWriter.Write(data)
	}

	return
}

// WriteHeader Fix: https://github.com/mholt/caddy/issues/38
func (g *handler) WriteHeader(code int) {
	g.Header().Del("Content-Length")
	g.ResponseWriter.WriteHeader(code)
}
