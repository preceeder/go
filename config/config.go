package config

import (
	"fmt"
	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/viper"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

type Config struct {
	Viper *viper.Viper
}

// NewConfigByFile 根据完整的文件名路径初始化配置
// cp string : base/config/cn.json
func NewConfigByFile(cp string) *Config {
	ConfigObj := &Config{Viper: viper.New()}
	ConfigObj.Viper.SetConfigFile(cp)
	ConfigObj.Viper.SetConfigType(path.Ext(cp)[1:])
	err := ConfigObj.Viper.ReadInConfig() // Find and read the config file
	if err != nil {
		fmt.Println(err.Error())
	}
	return ConfigObj
}

// NewConfigByMultiFile 根据路径初始化配置
// 注意各个配置文件内不要有重复的key, 不然会覆盖的
// cp string : base/config
func NewConfigByMultiFile(dir string) *Config {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(fmt.Errorf("read config dir failed: %w", err.Error()))
		os.Exit(1)
	}

	var configFiles []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(file.Name()))
		if ext == ".json" {
			configFiles = append(configFiles, filepath.Join(dir, file.Name()))
		}
	}
	sort.Strings(configFiles) // 可选：按字母顺序加载（保证一致性）
	v, err := LoadMultipleConfigs(configFiles)
	if err != nil {
		fmt.Println(fmt.Errorf("read config file failed: %w", err.Error()))
		os.Exit(1)
	}
	ConfigObj := &Config{Viper: v}
	return ConfigObj
}

func LoadMultipleConfigs(files []string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType("json")

	for i, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return nil, fmt.Errorf("open %s failed: %w", file, err)
		}
		defer f.Close()

		if i == 0 {
			// 第一个必须使用 ReadConfig
			if err := v.ReadConfig(f); err != nil {
				return nil, fmt.Errorf("read %s failed: %w", file, err)
			}
		} else {
			if err := v.MergeConfig(f); err != nil {
				return nil, fmt.Errorf("merge %s failed: %w", file, err)
			}
		}
	}
	return v, nil
}

/*
ReadViperConfig
将viper 配置读渠道 target 结构体中
*/
func ReadViperConfig(v viper.Viper, key string, target any) {
	err := v.UnmarshalKey(key, target, func(ms *mapstructure.DecoderConfig) { ms.TagName = "json" })
	if err != nil {
		fmt.Printf("load %s config error: %s\n", key, err.Error())
		os.Exit(1)
	}
	return
}
