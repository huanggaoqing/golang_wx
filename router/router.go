package router

import (
	"github.com/gin-gonic/gin"
)

func RegRouter() *gin.Engine {
	r := gin.Default()
	// 校验服务器路由
	r.GET("/", checkout)
	return r
}
