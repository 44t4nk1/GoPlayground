package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is the root page!",
		})
	})
	router.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Homepage!",
		})
	})
	router.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Login Page!",
		})
	})
	router.Run(":8000")
}
