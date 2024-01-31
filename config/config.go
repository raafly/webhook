package config

type AppConfig struct {
	App struct {
		Name string
	}
	Fiber struct {
		Port string
	}
	Postgres struct {
		Host     string
		Port     string
		Username string
		Name   	 string
		Password string
		SSL 	 string
	}
	Email struct {
		Sender	  string
		Adderss   string
		Password  string
	}
}

var appConfig *AppConfig

func NewAppConfig() *AppConfig {
	if appConfig == nil {
		appConfig = &AppConfig{}

		initConfig(appConfig)
	}
	return appConfig
}