package config

func initMail(conf *AppConfig) {
	v := initViper()

	conf.Email.Sender = v.GetString("EMAIL_SENDER_NAME")
	conf.Email.Adderss = v.GetString("EMAIL_SENDER_ADDRESS")
	conf.Email.Password = v.GetString("uhuvraenjuwewvkw")
}