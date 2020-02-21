package controller

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"github.com/treant5612/ytvc-web/service"
	"io"
	"log"
	"net/http"
	"strconv"
)

var (
	Page400 = `templates/400.html`
	Page500 = `templates/500.html`
)

func Video(c *gin.Context) {
	var err error
	urlParam := c.DefaultQuery("url", c.PostForm("url"))
	if urlParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"err":  "invalid param",
		})
		return
	}
	video, err := service.Video(urlParam)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"err":  "invalid param",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": video,
	})
}

func VideoDownload(c *gin.Context) {
	log.SetFlags(log.Lshortfile | log.Ltime)
	id := c.Param("id")
	noStr := c.Query("no")
	kind := c.Query("kind")
	no, err := strconv.Atoi(noStr)
	if err != nil {
		c.HTML(400, Page400, nil)
		return
	}
	fileName, downUrl, err := service.DownloadInfo(id, no, kind)
	if err != nil {
		c.HTML(404, Page400, nil)
		return
	}
	resp, err := service.Download(downUrl, func(r *http.Request) {
		copyHeader(r.Header, c.Request.Header, "Range")
	})
	if err != nil {
		log.Println(err)
		c.HTML(500, Page500, nil)
		return
	}
	defer resp.Body.Close()

	c.Status(resp.StatusCode)
	copyHeader(c.Writer.Header(), resp.Header,
		"Range", "Accept-Ranges", "Content-Range", "Content-Type", "Content-Length")
	c.Writer.Header().Set("Content-Disposition", "attachment;filename="+fileName)
	_, err = io.Copy(c.Writer, bufio.NewReaderSize(resp.Body, 512*1024))
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
