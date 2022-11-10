package router

import (
	"go_http/router/test"

	"github.com/gin-gonic/gin"
)

func RegRouter() *gin.Engine {
	r := gin.Default()
	test.RegTestRouter(r)
	return r
}
