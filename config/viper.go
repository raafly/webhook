package config

import (
	"github.com/spf13/viper"
)

func initViper() *viper.Viper {
	v := viper.New()
	v.SetConfigFile("/home/saturna/Desktop/Developments/GOLANG/src/webhook/.env")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return v
}
