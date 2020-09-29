package main

import (
	_ "gfast/boot"
	_ "gfast/router"
	"github.com/gogf/gf/frame/g"
)

// @title SealMS API文档
// @version 1.0
// @description SealMS 在线API文档
// @host localhost
// @BasePath /system
func main() {
	g.Server().Run()
}
