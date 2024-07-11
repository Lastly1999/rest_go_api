package global

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var GORM *gorm.DB
var Logger *zap.SugaredLogger
var Redis *redis.Client
