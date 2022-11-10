/*
 * @Author: huanggaoqing 3359793508@qq.com
 * @Date: 2022-09-08 19:59:30
 * @LastEditors: huanggaoqing 3359793508@qq.com
 * @LastEditTime: 2022-09-13 20:49:29
 * @FilePath: /go_http/router/test/router.go
 * @Description:
 *
 * Copyright (c) 2022 by huanggaoqing 3359793508@qq.com, All Rights Reserved.
 */
package test

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegTestRouter(r *gin.Engine) {
	test := r.Group("/test")
	{
		test.GET("/interface", func(c *gin.Context) {
			list := make([]int, 10)
			list = append(list, 6)
			fmt.Println(len(list))
			c.String(http.StatusOK, "test/interface")
		})
	}
}
