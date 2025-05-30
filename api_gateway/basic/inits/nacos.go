package inits

import (
	"api_gateway/basic/appconfig"
	"api_gateway/basic/global"
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go.uber.org/zap"
	"log"
)

func InitNacos() {
	clientConfig := constant.ClientConfig{
		NamespaceId:         appconfig.AppConf.NamespaceID, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      appconfig.AppConf.IdAddr,
			ContextPath: "/nacos",
			Port:        uint64(appconfig.AppConf.Port),
			Scheme:      "http",
		},
	}

	configClient, _ := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	content, _ := configClient.GetConfig(vo.ConfigParam{
		DataId: appconfig.AppConf.DataId,
		Group:  appconfig.AppConf.Group})
	log.Println(content)
	var my appconfig.Config
	json.Unmarshal([]byte(content), &my)
	global.Nacos = my
	log.Println(global.Nacos)

	err := configClient.ListenConfig(vo.ConfigParam{
		DataId: appconfig.AppConf.DataId,
		Group:  appconfig.AppConf.Group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)

			var newConfig appconfig.Config
			json.Unmarshal([]byte(data), &newConfig)
			global.Nacos = newConfig
			//log.Println("配置热更新成功:", global.Nacos)
			zap.L().Info("配置热更新成功")

		},
	})
	if err != nil {
		log.Fatalf("配置热更新失败")
	}
	//log.Println("Nacos配置监听已启动")
	zap.L().Info("nacos热更新成功")

}
