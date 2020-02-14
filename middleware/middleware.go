package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		host := context.Request.Host
		url := context.Request.URL
		method := context.Request.Method
		context.Next()
		status := context.Writer.Status()
		size := context.Writer.Size()
		log.Println(host, url, method, status, size)
	}
}
