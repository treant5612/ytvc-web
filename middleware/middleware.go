package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/treant5612/ytvc-web/db/sqldb"
	"github.com/treant5612/ytvc-web/model"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqTime := time.Now()
		c.Next()
		reqLog := &model.RequestLog{
			RemoteAddr:      c.ClientIP(),
			RequestUrl:      c.Request.RequestURI,
			RequestMethod:   c.Request.Method,
			RequestTime:     reqTime,
			RequestDuration: time.Since(reqTime),
			ResponseSize:    c.Writer.Size(),
			ResponseStatus:  c.Writer.Status(),
			Referer:         c.GetHeader("REFERER"),
		}
		sqldb.SaveRequestLog(reqLog)
	}
}

var limit_per_min = 10

func AccessControl() gin.HandlerFunc {
	ch := make(chan struct{}, limit_per_min)
	go func() {
		ticker := time.Tick(time.Minute / time.Duration(limit_per_min))
		for {
			select {
			case ch <- struct{}{}:
			default:
			}
			<-ticker
		}
	}()

	return func(c *gin.Context) {
		select {
		case <-ch:
			c.Next()
		default:
			c.JSON(429, gin.H{
				"error": "Too Many Requests",
			})
		}
	}
}
