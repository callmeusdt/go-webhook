package bootstrap

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	App      ConfigApp  `json:"app"`
	Commands []*Command `json:"commands"`
}

type ConfigApp struct {
	Name    string  `json:"name" yaml:"name"`
	Version float32 `json:"version" yaml:"version"`
	Port    string  `json:"port" yaml:"port"`
	Secret  string  `json:"secret" yaml:"secret"`
}

type Command struct {
	Name string   `json:"name" yaml:"name"`
	Args []string `json:"args" yaml:"args"`
}

func LoadConfig(configFile string) *Config {
	viper.SetConfigFile(fmt.Sprintf("%s.yaml", configFile))
	viper.AddConfigPath("../")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %w", err))
	}

	var _cfg *Config
	if err := viper.Unmarshal(&_cfg); err != nil {
		panic(fmt.Errorf("failed to unmarshal config: %w", err))
	}

	return _cfg
}
