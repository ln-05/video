package inits

import (
	"go.uber.org/zap"
	"log"
)

func Zap() {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"../api_gateway/zap.log"}
	build, err := config.Build()
	if err != nil {
		panic("日志初始化失败")
	}
	zap.ReplaceGlobals(build)
	log.Println("日志初始化成功")
}
