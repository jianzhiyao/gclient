package main

//导入包
import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mholt/archiver/v3"
	"net/http"
	_ "net/http/pprof"
	"strings"
	"time"
)

func any() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Header(`Content-Type`, `text/html`)
		context.String(http.StatusOK, `ok`)
	}
}

func main() {
	router := gin.Default()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	api := router.Use(
		gin.Recovery(),
	)
	{
		//UT
		ut := api.Use(
			compression(),
		)
		{
			ut.Any(`/`, any())
		}

		//benchmark
		api.Any(`/benchmark`, any())
	}

	s.ListenAndServe()
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

// Fix: https://github.com/mholt/caddy/issues/38
func (g *handler) WriteHeader(code int) {
	g.Header().Del("Content-Length")
	g.ResponseWriter.WriteHeader(code)
}
