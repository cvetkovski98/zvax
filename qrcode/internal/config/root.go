package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Db    DbConfig    `mapstructure:"db"`
	Pool  PoolConfig  `mapstructure:"db.pool"`
	Minio MinIOConfig `mapstructure:"minio"`
}

func LoadConfig(name string) error {
	viper.AddConfigPath("config")
	viper.SetConfigName(name)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	return viper.ReadInConfig()
}

func GetConfig() *Config {
	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		panic(err)
	}
	return cfg
}
