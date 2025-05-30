package main

import (
	"api_gateway/basic/inits"
	_ "api_gateway/basic/inits"
	"api_gateway/router"
	"github.com/gin-gonic/gin"
)

func main() {

	inits.Init()
	r := gin.Default()
	router.LoadRouter(r)
	r.Run(":8000") // 监听并在 0.0.0.0:8080 上启动服务
}
