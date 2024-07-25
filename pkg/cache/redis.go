package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"resetgoapi.com/go-admin-api/global"
	"resetgoapi.com/go-admin-api/pkg/config"
)

func Setup() {
	global.Logger.Info("Starting Redis setup...")
	redisConfig := config.GlobalConfig.Database.Redis
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host + ":" + redisConfig.Port,
		Password: redisConfig.Password, // 没有密码，默认值
		DB:       redisConfig.DB,       // 默认DB 0
	})

	global.Logger.Info("Pinging Redis server...")
	result := rdb.Ping(ctx)
	ok, err := result.Result()
	if err != nil {
		global.Logger.Error("Failed to ping Redis server. Error:", err)
		return
	}

	global.Logger.Info("Successfully pinged Redis server. Response:", ok)
	global.Redis = rdb
	global.Logger.Info("Redis setup completed successfully.")
}
