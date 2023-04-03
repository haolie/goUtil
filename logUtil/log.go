package logUtil

import (
	"go.uber.org/zap"
)

const (
	con_log_file    = "fileName"
	con_log_funName = "funName"
	con_log_content = "content"
	con_log_info    = "info"
)

var logger *zap.Logger

func InitLog() error {
	//var err error
	//logger, err = zap.NewDevelopment()
	//return err

	GetLogger()

	return nil
}

func InfoLog(info, fileName, funName, format string, args ...interface{}) {
	logger.Info(info, zap.String(con_log_file, fileName), zap.String(con_log_funName, funName), zap.String(con_log_content, format))
}

func ErrLog(fileName, funName string, format error) {

}

func DebugLog(format string, args ...interface{}) {

}

func FailLog(format string, args ...interface{}) {

}
