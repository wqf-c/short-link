package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"short-link/common"
	"short-link/service"
)

func main() {
	router := gin.Default()

	router.POST("/getShortLink", func(c *gin.Context) {
		url := c.PostForm("url")
		link, err := service.GetShortLink(url)
		if err != nil {
			c.JSON(200, gin.H{
				"errCode": 600001,
				"message": err.Error(),
			})
		} else {
			link = common.Info.Server.Ip + ":" + common.Info.Server.Port + "/" + link
			c.JSON(200, gin.H{
				"errCode": 0,
				"message": "",
				"link":    link,
			})
		}

	})

	router.GET("/:shortLink", func(c *gin.Context) {
		shortLink := c.Param("shortLink")
		link, err := service.GetLongLinkBySLink(shortLink)
		if err != nil {
			c.JSON(200, gin.H{
				"errCode": 600002,
				"message": err.Error(),
			})
		} else {
			c.Redirect(http.StatusMovedPermanently, link.LongLink)
		}

	})

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run(":" + common.Info.Server.Port)
}
