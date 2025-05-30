package inits

import "LiveStreaming_srv/basic/appconfig"

func Init() {
	appconfig.InitConfig()
	Zap()
	InitNacos()
	InitMysql()
	InitRedis()
}
