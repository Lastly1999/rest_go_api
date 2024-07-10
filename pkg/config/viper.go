package config

import (
	"github.com/spf13/viper"
	"resetgoapi.com/rest_go_api/global"
)

var GlobalConfig *AppConfig

func Setup() {
	// 读取 app.yaml 文件
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf") // 可以根据实际情况调整配置文件路径

	err := viper.ReadInConfig()
	if err != nil {
		global.Logger.Error("Error reading config file, %s", err)
	}

	// 获取 app.mode 的值
	mode := viper.GetString("app.mode")
	if mode == "" {
		global.Logger.Error("app.mode is not set in app.yaml")
	}

	// 根据 mode 加载相应的配置文件
	viper.SetConfigName(mode)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	err = viper.MergeInConfig()
	if err != nil {
		global.Logger.Error("Error reading %s config file, %s", mode, err)
	}
	err = viper.Unmarshal(&GlobalConfig)
	if err != nil {
		global.Logger.Error("Unable to decode into struct, %v", err)
	}
	global.Logger.Infof("Successfully loaded configuration for mode: %s", mode)
}
