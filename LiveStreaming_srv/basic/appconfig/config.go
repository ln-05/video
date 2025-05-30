package appconfig

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	MysqlConfig struct {
		User     string
		Password string
		Host     string
		Port     int
		Data     string
	}
	RedisConfig struct {
		Password string
		Host     string
		Db       int
	}
}

type Nacos struct {
	NamespaceID string
	DataId      string
	Group       string
	IdAddr      string
	Port        int
}

var AppConf Nacos

func InitConfig() {
	viper.SetConfigFile("../api_gateway/basic/appconfig/dev.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("配置文件读取失败")
	}
	log.Println("配置文件读取成功")
	err = viper.Unmarshal(&AppConf)
	if err != nil {
		panic("配置文件解析失败")
	}
	log.Println("配置文件解析成功")
}
