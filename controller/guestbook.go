package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/treant5612/ytvc-web/service"
	"log"
	"net/http"
)

func GuestBookPage(c *gin.Context) {
	pageNo := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("size", "15")

	comments, err := service.ListComment(pageNo, pageSize)
	if err != nil {
		log.Println("list comments failed", err)
	}
	c.HTML(http.StatusOK, "templates/guestbook.html", comments)
}

func GuestBookComment(c *gin.Context) {
	nickname := c.PostForm("nickname")
	content := c.PostForm("comment")
	log.Println("n", nickname, "c", content)
	if nickname != "" && content != "" {
		err := service.InsertComment(nickname, content, c.Request.RemoteAddr)
		if err == nil {
			c.JSON(200, gin.H{
				"msg": "comment success",
			})
			return
		}
		log.Println(err)
	}

	c.JSON(400, gin.H{
		"msg": "invalid parameter",
	})

}
