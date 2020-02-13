package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/treant5612/ytvc-web/service"
	"log"
)

func CaptionDownload(c *gin.Context) {
	captionId := c.Param("id")
	fname := c.DefaultQuery("fname", "subtitle")
	tlang := c.Query("tlang")
	secondaryId := c.Query("secondary")
	secondaryTlang := c.Query("secondary_tlang")
	var path string
	var err error
	if secondaryId == "" {
		path, err = service.DownloadCaption(captionId, tlang)
	} else {
		path, err = service.DownloadAndMergeCaption(captionId, tlang, secondaryId, secondaryTlang)
	}
	if err != nil {
		log.Println(err)
		c.Status(500)
		return
	}
	//	defer os.Remove(path)
	c.FileAttachment(path, fmt.Sprintf("%s.srt", fname))
}
