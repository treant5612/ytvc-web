package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/treant5612/ytvc-web/service"
	"io"
	"log"
	"net/http"
	"net/url"
)

func Video(c *gin.Context) {
	var err error
	var message string
	defer func() {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"err":     err.Error(),
				"message": message,
			})
		}
	}()
	url := c.Query("url")
	info, err := service.VideoInfo(url)
	if err != nil {
		message = fmt.Sprintf("Get video info failed:%v", err)
		return
	}
	captions, err := service.Captions(info.ID)
	if err != nil {
		log.Println(err)
	} else {
		info.Captions = captions
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": info,
	})
}

func VideoDownload(c *gin.Context) {
	log.SetFlags(log.Lshortfile | log.Ltime)
	downUrl := c.Query("url")
	if _, err := url.Parse(downUrl); err != nil {
		c.HTML(400, "index.html", nil)
		return
	}

	resp, err := service.Download(downUrl, func(r *http.Request) {
		copyHeader(r.Header, c.Request.Header, "Range")
	})
	if err != nil {
		log.Println(err)
		c.HTML(500, "index.html", nil)
		return

	}
	defer resp.Body.Close()
	c.Status(resp.StatusCode)
	copyHeader(c.Writer.Header(), resp.Header,
		"Range", "Accept-Ranges", "Content-Range", "Content-Type", "Content-Length")

	//c.Writer.Header().Set("Content-Disposition", "attachment;filename=test.mp4")

	_, err = io.Copy(c.Writer, resp.Body)
	if err != nil {
		log.Println(err)
	}

}

func copyHeader(dst, src http.Header, fields ...string) {
	for _, field := range fields {
		if v := src.Get(field); v != "" {
			dst.Set(field, v)
		}
	}
}
