package main

import (
	"go_http/router"
)

func main() {
	r := router.RegRouter()
	r.Run(":80")
}
