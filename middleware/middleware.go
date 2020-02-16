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
			RemoteAddr:      c.Request.RemoteAddr,
			RequestUrl:      c.Request.RequestURI,
			RequestMethod:   c.Request.Method,
			RequestTime:     reqTime,
			RequestDuration: time.Since(reqTime),
			ResponseSize:    c.Writer.Size(),
			ResponseStatus:  c.Writer.Status(),
		}
		sqldb.SaveRequestLog(reqLog)
	}
}

func AccessControl() gin.HandlerFunc {
	return func(c *gin.Context) {
		//todo
		if false {
			c.HTML(400, "templates/400.html", nil)
			return
		}

		c.Next()
	}
}
