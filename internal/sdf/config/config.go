package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Profile string
}

func InitConfig() *Config {
	viper.SetConfigType("yml")
	viper.SetDefault("projectdir", "$HOME/src")
	viper.SetDefault("profile", "main")

	viper.AddConfigPath("$HOME/.sdf")
	viper.AddConfigPath(".")

	return &Config{}
}

func (c *Config) ReadConfig(configName string) {
	viper.SetConfigName(configName)
	viper.Set("profile", configName)
	viper.ReadInConfig()

	c.applyInterpolation()
}

func (c *Config) applyInterpolation() {
	viper.Set(
		"projectdir",
		os.ExpandEnv(viper.GetString("projectdir")),
	)
}
