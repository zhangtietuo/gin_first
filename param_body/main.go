package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {

	r := gin.Default()
	r.POST("/testBody", func(c *gin.Context) {
		bodyByts, err:=ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
		}
		//之前body中数据已经取出到bodyByts 现在需要将bodyByts回传入Body 才能获取数据
		c.Request.Body=ioutil.NopCloser(bytes.NewBuffer(bodyByts))
		firstName := c.PostForm("first_name'")
		lastName := c.DefaultPostForm("last_name", "default_last_name")

		c.String(http.StatusOK,"%s,%s,%s", firstName, lastName, string(bodyByts))
	})


	r.Run(":8080")
}