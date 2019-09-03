package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/:ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": c.Param("ping"),
			"fuck": 123,
		})
	})

	//泛绑定 /user/为前缀的url 都能接收到
	r.GET("/user/*action", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "adssad",
			"fuck": 123,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}