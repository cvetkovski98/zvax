package config

import "fmt"

type MinIOConfig struct {
	Endpoint     string `mapstructure:"endpoint"`
	RootUser     string `mapstructure:"root_user"`
	RootPassword string `mapstructure:"root_password"`
	UseSSL       bool   `mapstructure:"use_ssl"`
	BucketName   string `mapstructure:"bucket_name"`
	Location     string `mapstructure:"location"`
}

func (c *MinIOConfig) Dsn() string {
	return fmt.Sprintf("%s://%s:%s@%s", "https", c.RootUser, c.RootPassword, c.Endpoint)
}

func (c *MinIOConfig) ResourceLocation() string {
	return fmt.Sprintf("%s/%s", c.Endpoint, c.BucketName)
}
