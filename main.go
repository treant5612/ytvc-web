package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/treant5612/ytvc-web/controller"
	"github.com/treant5612/ytvc-web/db/redisdb"
	"github.com/treant5612/ytvc-web/db/sqldb"
	"github.com/treant5612/ytvc-web/middleware"
	"github.com/treant5612/ytvc-web/service"
	"log"
)

func main() {
	prepare()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(middleware.Logger())
	router.LoadHTMLGlob("templates/*")
	y2b := router
	{
		y2b.Static("/resources", "./resources")
		y2b.StaticFile("/favicon.ico", "./resources/favicon.ico")
		y2b.GET("/", controller.Index)
		y2b.GET("/index.html", controller.Index)

		y2b.GET("/video", controller.Video)
		y2b.POST("/video", controller.Video)
		y2b.GET("/video/:id", controller.VideoDownload)
		y2b.GET("/caption/:id", controller.CaptionDownload)

		y2b.GET("/guestbook", controller.GuestBookPage)
		y2b.POST("/guestbook", controller.GuestBookComment)
	}
	router.Run(":8080")
}

func prepare() {
	log.SetFlags(log.Lshortfile | log.Ltime)
	//youtubeapi.InitServiceFSC("client_secret.json", "youtubeForceSslToken.json")
	redisdb.Init(&redis.Options{Addr: ":6379"})
	sqldb.InitSqlite("./sql.db")
	service.SetDownloadPath("/dev/shm")

}
