package env

import (
	"fmt"
	"github.com/spf13/viper"
)

type config struct {
	DBDSN    string `mapstructure:"DB_DSN"`
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBPool   int    `mapstructure:"DB_POOL"`
}

func NewEnv(path string) (*config, error) {
	fmt.Println("path", path)
	var cfg *config
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
