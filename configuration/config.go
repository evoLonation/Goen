package configuration

import "github.com/spf13/viper"

type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	//添加配置文件路径
	viper.AddConfigPath(path)
	//配置文件名
	viper.SetConfigName("app")
	//配置文件后缀名
	viper.SetConfigType("env")

	// 从环境变量读取可能存在的对应文件中的值，如果存在就覆盖
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	// 将读取到的绑定到结构体中
	err = viper.Unmarshal(&config)

	return

}
