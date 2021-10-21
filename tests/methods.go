package tests

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ok() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Header(`Content-Type`, `text/html`)
		context.String(http.StatusOK, `ok`)
	}
}

