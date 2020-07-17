package router

import (
	"net/http"

	"github.com/chen-shiwei/shorturl/internal/shorturl/handler"
	"github.com/chen-shiwei/shorturl/pkg/long2short"
	"github.com/gin-gonic/gin"
)

func New() http.Handler {
	r := gin.Default()
	r.GET("/link/:short", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, long2short.GetShortMap(c.Param("short")))
	})
	r.POST("long2short", handler.Long2Short())
	return r
}
