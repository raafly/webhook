package config

func initApp(conf *AppConfig) {
	v := initViper()

	conf.App.Name = v.GetString("APP_NAME")
}