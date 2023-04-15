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

var CFG Config

func LoadConfig() {
	viper.SetConfigFile("main.yaml")
	viper.AddConfigPath("../")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %w", err))
	}

	if err := viper.Unmarshal(&CFG); err != nil {
		panic(fmt.Errorf("failed to unmarshal config: %w", err))
	}
	fmt.Printf("Loaded config: %+v\n", CFG)
}
