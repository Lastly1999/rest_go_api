package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func Setup() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	Logger = sugar
}
