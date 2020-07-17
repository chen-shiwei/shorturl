package handler

import (
	"net/http"

	"github.com/chen-shiwei/shorturl/pkg/long2short"
	"github.com/gin-gonic/gin"
)

type Long2ShortArgs struct {
	Url string `json:"url"`
}

func Long2Short() gin.HandlerFunc {
	return func(c *gin.Context) {

		var args Long2ShortArgs

		err := c.ShouldBind(&args)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
			return
		}

		short, err := long2short.LongUrl2ShortUrl(args.Url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "服务不可用"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": gin.H{
			"short_url": c.Request.Host + "/link/" + short,
			"long_url":  args.Url,
		}, "message": "success"})
		return
	}
}
