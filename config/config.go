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
		Password string
		Dbname   string
	}
}

var appConfig *AppConfig

func NewAppConfig() *AppConfig {
	if appConfig == nil {
		appConfig = &AppConfig{}

		initApp(appConfig)
		initFiber(appConfig)
		initPostgres(appConfig)
	}
	return appConfig
}
