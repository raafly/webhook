package database

import (
	"fmt"
	"log"

	"github.com/raafly/webhook/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(conf *config.AppConfig) *gorm.DB {
	host := conf.Postgres.Host
	port := conf.Postgres.Port
	user := conf.Postgres.Username
	pass := conf.Postgres.Password
	name := conf.Postgres.Dbname

	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, pass, name)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
