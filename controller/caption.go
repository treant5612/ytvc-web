package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/treant5612/ytvc-web/service"
	"log"
)

func CaptionDownload(c *gin.Context) {
	videoId := c.Param("id")
	captionId := c.Query("vssid")
	fname := c.DefaultQuery("fname", "subtitle")
	tlang := c.Query("tlang")
	secondaryId := c.Query("secondary")
	secondaryTlang := c.Query("secondary_tlang")
	var fpath string
	var err error
	if secondaryId == "" {
		fpath, err = service.DownloadCaption(videoId, captionId, tlang)
	} else {
		fpath, err = service.DownloadAndMergeCaption(videoId, captionId, tlang, secondaryId, secondaryTlang)
	}
	if err != nil {
		log.Println(err)
		c.Status(500)
		return
	}
	//	defer os.Remove(fpath)
	fileName := fmt.Sprintf("%s.srt", fname)
	c.FileAttachment(fpath, fileName)
}
