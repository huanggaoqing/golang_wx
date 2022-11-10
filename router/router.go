package router

import (
	"github.com/gin-gonic/gin"
)

func RegRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", checkout)
	return r
}
