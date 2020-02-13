package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/treant5612/ytvc-web/service"
	"github.com/treant5612/ytvc-web/utils"
	"io"
	"log"
	"net/http"
	"strconv"
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
	urlParam := c.Query("url")
	video, err := service.Video(urlParam)
	if err != nil {
		log.Println(err)
	}

	//将字幕语言标识符转换为中文显示
	for i := range video.Captions {
		video.Captions[i].Snippet.Language = utils.LanguageDisplay(video.Captions[i].Snippet.Language)
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
	log.Println(id, noStr)
	no, err := strconv.Atoi(noStr)

	if err != nil {
		c.HTML(400, "index.html", nil)
		return
	}
	fileName, downUrl, err := service.DownloadInfo(id, no)
	if err != nil {
		c.HTML(404, "index.html", nil)
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
	c.Writer.Header().Set("Content-Disposition", "attachment;filename="+fileName)
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
