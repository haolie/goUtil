package logUtil

import (
	"fmt"

	"go.uber.org/zap"
)

const (
	con_log_file    = "fileName"
	con_log_funName = "funName"
	con_log_content = "content"
	con_log_info    = "info"
)

var logger *zap.Logger

func InitLog() {
	GetLogger()
}

func InfoLog(info, fileName, funName, format string, args ...interface{}) {
	if len(args) > 0 {
		format = fmt.Sprintf(format, args)
	}
	logger.Info(info, zap.String(con_log_file, fileName), zap.String(con_log_funName, funName), zap.String(con_log_content, format))
}

func ErrLog(info string, err error) {
	logger.Error(info, zap.String("Err", fmt.Sprintf("%v", err)))
}

func DebugLog(info, format string, args ...interface{}) {
	if len(args) > 0 {
		format = fmt.Sprintf(format, args)
	}

	logger.Debug(info, zap.String("debug", format))
}

func FailLog(format string, args ...interface{}) {
	if len(args) > 0 {
		format = fmt.Sprintf(format, args)
	}
	logger.Fatal("fail", zap.String("fail", format))
}
