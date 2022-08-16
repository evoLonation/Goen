package configuration

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

var Conf *Config

type Config struct {
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func init() {
	var err error
	Conf, err = loadConfig("./")
	if err != nil {
		log.Fatal(err)
	}
}

func loadConfig(path string) (*Config, error) {
	//添加配置文件路径
	viper.AddConfigPath(path)
	//配置文件名
	viper.SetConfigName("app")
	//配置文件后缀名
	viper.SetConfigType("env")

	// 从环境变量读取可能存在的对应文件中的值，如果存在就覆盖
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Config{}
	// 将读取到的绑定到结构体中
	err := viper.Unmarshal(config)
	return config, err
}
