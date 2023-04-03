module github.com/haolie/goUtil

go 1.17

replace gopkg.in/natefinch/lumberjack.v2 v2.2.1 => github.com/natefinch/lumberjack v0.0.0

replace github.com/natefinch/lumberjack v0.0.0 => gopkg.in/natefinch/lumberjack.v2 v2.2.1

require (
	github.com/lestrrat/go-file-rotatelogs v0.0.0-20180223000712-d3151e2a480f
	github.com/natefinch/lumberjack v0.0.0
	go.uber.org/zap v1.24.0
)

require (
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/jonboulle/clockwork v0.4.0 // indirect
	github.com/lestrrat/go-envload v0.0.0-20180220120943-6ed08b54a570 // indirect
	github.com/lestrrat/go-strftime v0.0.0-20180220042222-ba3bf9c1d042 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/tebeka/strftime v0.1.5 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/net v0.8.0 // indirect
)
