module github.com/haolie/goUtil

go 1.17

replace gopkg.in/natefinch/lumberjack.v2 v2.2.1 => github.com/natefinch/lumberjack v0.0.0

replace github.com/natefinch/lumberjack v0.0.0 => gopkg.in/natefinch/lumberjack.v2 v2.2.1

require (
	github.com/natefinch/lumberjack v0.0.0
	go.uber.org/zap v1.24.0
)

require (
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
)
