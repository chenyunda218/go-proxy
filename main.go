package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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
		url := fmt.Sprintf("https://%s/?wc-api=%s", notifyUrl, uri)
		jsonStr, _ := json.Marshal(&body)
		req, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(string(jsonStr)))
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
