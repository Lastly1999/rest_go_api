package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

var Config AppConfig

func Setup() {
	config := "app.yaml"
	if configEnv := os.Getenv("APP_ENV"); configEnv != "" {
		config = configEnv
	}
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config fileDriver changed:", in.Name)
		// 重载配置,这里可以进行重启服务器,或者其他操作
		if err := v.Unmarshal(&Config); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&Config); err != nil {
		fmt.Println(err)
	}
}
