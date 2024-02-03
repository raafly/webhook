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

func initConfig(conf *AppConfig) {
	v := initViper()

	conf.App.Name = v.GetString("APP_NAME")

	conf.Fiber.Port = v.GetString("FIBER_PORT")

	conf.Email.Sender = v.GetString("EMAIL_SENDER_NAME")
	conf.Email.Adderss = v.GetString("EMAIL_SENDER_ADDRESS")
	conf.Email.Password = v.GetString("EMAIL_SENDER_PASSWORD")

	conf.Postgres.Host = v.GetString("DB_HOST")
	conf.Postgres.Port = v.GetString("DB_PORT")
	conf.Postgres.Username = v.GetString("DB_USER")
	conf.Postgres.Password = v.GetString("DB_PASSWORD")
	conf.Postgres.Name = v.GetString("DB_NAME")
	conf.Postgres.SSL = v.GetString("DB_SSL")
}

// func initApp(conf *AppConfig) {
// 	v := initViper()

// 	conf.App.Name = v.GetString("APP_NAME")
// }

// func initFiber(conf *AppConfig) {
// 	v := initViper()

// 	conf.Fiber.Port = v.GetString("FIBER_PORT")
// }

// func initMail(conf *AppConfig) {
// 	v := initViper()

// 	conf.Email.Sender = v.GetString("EMAIL_SENDER_NAME")
// 	conf.Email.Adderss = v.GetString("EMAIL_SENDER_ADDRESS")
// 	conf.Email.Password = v.GetString("uhuvraenjuwewvkw")
// }

// func initPostgres(conf *AppConfig) {
// 	v := initViper()

// 	conf.Postgres.Host = v.GetString("DB_HOST")
// 	conf.Postgres.Port = v.GetString("DB_PORT")
// 	conf.Postgres.Username = v.GetString("DB_USER")
// 	conf.Postgres.Password = v.GetString("DB_PASSWORD")
// 	conf.Postgres.Name = v.GetString("DB_NAME")
// 	conf.Postgres.SSL = v.GetString("DB_SSL")
// }