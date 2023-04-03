package logUtil

import (
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	logTmFmt            = "2006-01-02 15:04:05"
	con_default_logPath = "./log"
)

//
func GetLogger() {
	Encoder := GetEncoder()
	WriteSyncer := GetWriteSyncer()
	LevelEnabler := GetLevelEnabler()
	ConsoleEncoder := GetConsoleEncoder()
	newCore := zapcore.NewTee(
		zapcore.NewCore(Encoder, WriteSyncer, LevelEnabler),                          // 写入文件
		zapcore.NewCore(ConsoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel), // 写入控制台
	)
	logger = zap.New(newCore, zap.AddCaller())
	zap.ReplaceGlobals(logger)
}

// GetEncoder 自定义的Encoder
func GetEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(
		zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller_line",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     "\r\n",
			EncodeLevel:    cEncodeLevel,
			EncodeTime:     cEncodeTime,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   cEncodeCaller,
		})
}

// GetConsoleEncoder 输出日志到控制台
func GetConsoleEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
}

// GetWriteSyncer 自定义的WriteSyncer
func GetWriteSyncer() zapcore.WriteSyncer {
	//lumberJackLogger := &lumberjack.Logger{
	//	Filename:   "./zap.log",
	//	MaxSize:    200,
	//	MaxBackups: 10,
	//	MaxAge:     30,
	//}
	return zapcore.AddSync(getWriter())
	//return zapcore.AddSync(lumberJackLogger)
}

func getWriter() io.Writer {
	f, err := os.Stat(con_default_logPath)
	if err != nil || f.IsDir() == false {
		err = os.MkdirAll(con_default_logPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	hook, err := rotatelogs.New(
		con_default_logPath+"/%Y-%m-%d_%H.log",
		rotatelogs.WithLinkName(con_default_logPath),
		rotatelogs.WithMaxAge(time.Hour*24*30),
		rotatelogs.WithRotationTime(time.Hour))

	if err != nil {
		panic(err)
	}

	return hook
}

// GetLevelEnabler 自定义的LevelEnabler
func GetLevelEnabler() zapcore.Level {
	return zapcore.InfoLevel
}

// cEncodeLevel 自定义日志级别显示
func cEncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

// cEncodeTime 自定义时间格式显示
func cEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + t.Format(logTmFmt) + "]")
}

// cEncodeCaller 自定义行号显示
func cEncodeCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + caller.TrimmedPath() + "]")
}
