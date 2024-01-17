package config

func initFiber(conf *AppConfig) {
	v := initViper()

	conf.Fiber.Port = v.GetString("FIBER_PORT")
}
