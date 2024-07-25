package models

import "resetgoapi.com/go-admin-api/models"

type SysConfig struct {
	models.BaseModel
	// 参数名称
	ConfigName string `json:"configName" form:"configName"`
	// 参数键名
	ConfigKey string `json:"configKey" form:"configKey"`
	// 参数键值
	ConfigValue string `json:"configValue" form:"configValue"`
	// 系统内置 0 否 1 是
	ConfigType bool `json:"configType" form:"configType"`
	// 备注
	Remark string `json:"remark" form:"remark"`
}
