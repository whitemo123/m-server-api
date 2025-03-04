package config

import (
	"bytes"
	_ "embed"
	"io"
	"m-server-api/pkg/env"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Log      LogConfig      `yaml:"log"`
	Database DatabaseConfig `yaml:"database"`
	Jwt      JwtConfig      `yaml:"jwt"`
}

type ServerConfig struct {
	Port   int      `yaml:"port"`
	Mode   string   `yaml:"mode"`
	Prefix string   `yaml:"prefix"`
	White  []string `yaml:"white"`
}

type LogConfig struct {
	Level string `yaml:"level"`
}

type DatabaseConfig struct {
	MySQL MySQLConfig `yaml:"mysql"`
	Redis RedisConfig `yaml:"redis"`
}

type MySQLConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

type JwtConfig struct {
	Secret string `yaml:"secret"`
	Expire int    `yaml:"expire"`
}

var config = new(Config)

var (
	//go:embed dev.yaml
	devConfig []byte

	//go:embed test.yaml
	testConfig []byte

	//go:embed prod.yaml
	prodConfig []byte
)

func init() {
	var r io.Reader

	switch env.Active().Value() {
	case "dev":
		r = bytes.NewReader(devConfig)
	case "test":
		r = bytes.NewReader(testConfig)
	case "prod":
		r = bytes.NewReader(prodConfig)
	default:
		r = bytes.NewReader(devConfig)
	}

	viper.SetConfigType("yaml")

	if err := viper.ReadConfig(r); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}
}

func Get() Config {
	return *config
}
