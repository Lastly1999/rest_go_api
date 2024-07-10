package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var GORM *gorm.DB
var Logger *zap.SugaredLogger
