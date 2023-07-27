package main

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/woocommerce-api/:notifyUrl/:uri", func(c *gin.Context) {
		notifyUrl := c.Param("notifyUrl")
		uri := c.Param("uri")
		var body map[string]interface{}
		c.ShouldBindJSON(&body)
		client := &http.Client{}
		url := fmt.Sprintf("%s/?wc-api=%s", notifyUrl, uri)
		fmt.Println(url)
		req, _ := http.NewRequest(http.MethodPost, url, nil)
		resp, err := client.Do(req)
		if err != nil || resp.StatusCode == 404 {
			c.String(404, "error")
			return
		}
	})
	router.GET("/template", func(ctx *gin.Context) {
		base64url := ctx.Query("html")
		b, err := base64.URLEncoding.DecodeString(base64url)
		if err != nil {
			fmt.Println("err")
		}
		ctx.Writer.Header().Set("Content-Type", "text/html; charset=Big5")
		ctx.String(200, string(b))
	})
	router.Run(":10000")
}
