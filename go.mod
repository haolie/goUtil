module github.com/haolie/goUtil

go 1.23.0

//toolchain go1.23.6

replace gopkg.in/natefinch/lumberjack.v2 v2.2.1 => github.com/natefinch/lumberjack v0.0.0

replace github.com/natefinch/lumberjack v0.0.0 => gopkg.in/natefinch/lumberjack.v2 v2.2.1

require (
	github.com/lestrrat/go-file-rotatelogs v0.0.0-20180223000712-d3151e2a480f
	github.com/natefinch/lumberjack v0.0.0
	go.uber.org/zap v1.24.0
)

require (
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/go-viper/mapstructure/v2 v2.4.0 // indirect
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/jonboulle/clockwork v0.4.0 // indirect
	github.com/lestrrat/go-envload v0.0.0-20180220120943-6ed08b54a570 // indirect
	github.com/lestrrat/go-strftime v0.0.0-20180220042222-ba3bf9c1d042 // indirect
	github.com/pelletier/go-toml/v2 v2.2.4 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/sagikazarmark/locafero v0.11.0 // indirect
	github.com/sourcegraph/conc v0.3.1-0.20240121214520-5f936abd7ae8 // indirect
	github.com/spf13/afero v1.15.0 // indirect
	github.com/spf13/cast v1.10.0 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	github.com/spf13/viper v1.21.0 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/tebeka/strftime v0.1.5 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.yaml.in/yaml/v3 v3.0.4 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.28.0 // indirect
)
