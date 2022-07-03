package configs

import (
	"github.com/spf13/viper"
)

func NewConfigProvider() (error, *viper.Viper) {
	v := viper.New()

	v.SetConfigName("app")
	v.AddConfigPath("./conf")

	err := v.ReadInConfig()
	if err != nil {
		return err, nil
	}
	return nil, v
}
