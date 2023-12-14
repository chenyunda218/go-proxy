package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/woocommerce-api/:notifyUrl/:uri", func(c *gin.Context) {
		notifyUrl := c.Param("notifyUrl")
		uri := c.Param("uri")
		client := &http.Client{}
		url := fmt.Sprintf("https://%s/?wc-api=%s", notifyUrl, uri)
		req, _ := http.NewRequest(http.MethodPost, url, c.Request.Body)
		resp, err := client.Do(req)
		if err != nil || resp.StatusCode == 404 {
			c.String(404, "error")
			return
		} else {
			c.String(200, "success")
		}
	})
	router.Run(":10000")
}
