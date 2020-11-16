package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hu shewom chu ane hu lodu chu",
		})
	})
	router.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "homepage",
		})
	})
	router.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "login",
		})
	})
	router.Run(":8000")
}
