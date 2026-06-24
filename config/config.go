package config

import "github.com/spf13/viper"

const configFileName = "Config"

var (
	configMap      = make(map[string]interface{}, 8)
	isLoadFinished = false
)

// AddConfig 在 Load 完成前追加配置项。Load 之后调用将 panic。
func AddConfig(k string, v interface{}) {
	if isLoadFinished {
		panic("config already loaded")
	}
	configMap[k] = v
}

// UpdateConfig 更新已存在的配置项，key 不存在时不做任何操作。
func UpdateConfig(k string, v interface{}) {
	if _, exists := configMap[k]; exists {
		configMap[k] = v
	}
}

// GetValue 读取指定 key 的配置值并转换为类型 T。key 不存在时 exists 为 false。
// 类型不匹配时会 panic，调用方需确保类型与 toml 中一致。
func GetValue[T int32 | int64 | string | bool](k string) (v T, exists bool) {
	temp, exists := configMap[k]
	if exists {
		v = temp.(T)
	}
	return
}

// Load 从指定目录读取 Config.toml 并将所有配置项写入内存。
// path 为空时默认读取当前目录。Load 成功后 isLoadFinished 置为 true，禁止继续调用 AddConfig。
func Load(path string) error {
	viper.SetConfigName(configFileName)
	viper.SetConfigType("toml")
	if len(path) == 0 {
		path = "."
	}
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	for k, v := range viper.AllSettings() {
		configMap[k] = v
	}

	return nil
}

func LoadComplete() {
	isLoadFinished = true
}
