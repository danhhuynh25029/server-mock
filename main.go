package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

func main() {
	r := gin.Default()
	r.GET("/:path", func(c *gin.Context) {
		path := c.Request.URL.String()

		parse, err := url.Parse(path)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": err,
			})
			return
		}
		for _, v := range parse.Query() {
			fmt.Println(v)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": path,
		})
	})
	r.Run()
}
