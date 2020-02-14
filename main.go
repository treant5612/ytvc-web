package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/treant5612/ytvc-web/controller"
	"github.com/treant5612/ytvc-web/db/redisdb"
	"github.com/treant5612/ytvc-web/manager/youtubeapi"
	"github.com/treant5612/ytvc-web/middleware"
	"github.com/treant5612/ytvc-web/service"
	"log"
)

func main() {
	prepare()

	router := gin.Default()
	router.Use(middleware.Logger())
	router.LoadHTMLGlob("templates/*")
	y2b := router
	{
		y2b.Static("/resources", "./resources")
		y2b.Static("/page", "./static")
		y2b.GET("/", controller.Index)
		y2b.GET("/index.html", controller.Index)
		y2b.GET("/test/test/test", controller.Index)
		y2b.GET("/video", controller.Video)
		y2b.GET("/video/:id", controller.VideoDownload)
		y2b.GET("/caption/:id", controller.CaptionDownload)
	}
	router.Run("localhost:8080")
}

func prepare() {
	log.SetFlags(log.Lshortfile | log.Ltime)
	youtubeapi.InitServiceFSC("client_secret.json", "youtubeForceSslToken.json")
	redisdb.Init(&redis.Options{Addr: ":6379"})
	service.SetDownloadPath("/dev/shm")

}
