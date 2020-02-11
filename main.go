package main

import (
	"github.com/gin-gonic/gin"
	"github.com/treant5612/ytvc-web/controller"
	"github.com/treant5612/ytvc-web/manage/youtubeapi"
)

func main() {
	prepare()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	y2b := router.Group("/y2b")
	{
		y2b.Static("/resources", "./resources")
		y2b.GET("/", controller.Index)
		y2b.GET("/index.html", controller.Index)
		y2b.GET("/test/test/test", controller.Index)
		y2b.GET("/video", controller.Video)
		y2b.GET("/video/dl", controller.VideoDownload)
	}

	router.Run("localhost:8080")
}

func prepare() {
	youtubeapi.InitServiceFSC("client_secret.json", "youtubeForceSslToken.json")
}
