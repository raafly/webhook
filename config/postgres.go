package config

func initPostgres(conf *AppConfig) {
	v := initViper()

	conf.Postgres.Host = v.GetString("DB_HOST")
	conf.Postgres.Port = v.GetString("DB_PORT")
	conf.Postgres.Username = v.GetString("DB_USER")
	conf.Postgres.Password = v.GetString("DB_PASSWORD")
	conf.Postgres.Name = v.GetString("DB_NAME")
	conf.Postgres.SSL = v.GetString("DB_SSL")
}