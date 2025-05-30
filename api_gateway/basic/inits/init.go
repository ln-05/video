package inits

import "api_gateway/basic/appconfig"

func Init() {
	appconfig.InitConfig()
	InitNacos()
	InitMysql()
	InitRedis()
	Zap()
}
