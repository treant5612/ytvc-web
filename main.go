package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.Run("localhost:8080")
	router.Static("/resources/", "./resources")

	router.Run("localhost:8080")
}
