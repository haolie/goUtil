package logutil

import (
	"log/slog"
	"os"
)

var (
	debugLogger *slog.Logger
	infoLogger  *slog.Logger
	warnLogger  *slog.Logger
	errorLogger *slog.Logger
)

type logLvl int

// Level 实现 slog.Leveler 接口，返回自定义级别（低于 Debug，用于控制台全量输出）
func (lvl logLvl) Level() slog.Level {
	return -10
}

// InitLog 初始化日志系统。basePath 为日志文件根目录，按小时轮转写入文件，同时输出到控制台。
func InitLog(basePath string) {
	debugLogger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
		Level:     logLvl(0),
	}))

	infoLogger = slog.New(slog.NewTextHandler(newLogWriter(basePath, "info"), nil))
	warnLogger = slog.New(slog.NewTextHandler(newLogWriter(basePath, "warn"), nil))
	errorLogger = slog.New(slog.NewTextHandler(newLogWriter(basePath, "err"), nil))
}

// Debug 输出调试日志，仅写入控制台。
func Debug(str string) {
	debugLogger.Debug(str)
}

// Info 输出信息日志，写入文件（info）并同步输出到控制台。
func Info(str string) {
	infoLogger.Info(str)
	debugLogger.Info(str)
}

// Warn 输出警告日志，写入文件（warn）并同步输出到控制台。
func Warn(str string) {
	warnLogger.Warn(str)
	debugLogger.Warn(str)
}

// Error 输出错误日志，写入文件（err）并同步输出到控制台。
func Error(str string) {
	errorLogger.Error(str)
	debugLogger.Error(str)
}
