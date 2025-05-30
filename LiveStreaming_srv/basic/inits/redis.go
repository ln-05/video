package inits

import (
	"LiveStreaming_srv/basic/global"
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func InitRedis() {
	global.Rdb = redis.NewClient(&redis.Options{
		Addr:     global.Nacos.RedisConfig.Host,
		Password: global.Nacos.RedisConfig.Password,
		DB:       global.Nacos.RedisConfig.Db,
	})
	err := global.Rdb.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
		return
	}
	zap.L().Info("redis连接成功")

}
