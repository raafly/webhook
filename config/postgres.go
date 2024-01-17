package config

func initPostgres(conf *AppConfig) {
	v := initViper()

	conf.Postgres.Host = v.GetString("POSTGRES_HOST")
	conf.Postgres.Port = v.GetString("POSTGRES_PORT")
	conf.Postgres.Username = v.GetString("POSTGRES_USERNAME")
	conf.Postgres.Password = v.GetString("POSTGRES_PASSWORD")
	conf.Postgres.Dbname = v.GetString("POSTGRES_DBNAME")
}
