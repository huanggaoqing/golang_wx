package main

import (
	"go_http/router"
	"go_http/utils"
)

func main() {
	utils.GetEvendaysen()
	r := router.RegRouter()
	r.Run(":80")
}
