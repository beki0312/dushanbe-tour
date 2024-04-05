package db

import (
	"fmt"
	"go.uber.org/fx"
	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var NewPostgres = fx.Provide(newPostgres)

type dependencies struct {
	fx.In
}

type postgres struct {
	Postgres *gorm.DB
}

type IPostgres interface {
	GetPostgresConnection() *gorm.DB
}

func newPostgres(dp dependencies) IPostgres {
	host := "localhost"
	port := "5432"
	user := "app"
	dbname := "db"
	password := "pass"
	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Dushanbe", host, user, password, dbname, port)
	conn, err := gorm.Open(postgresDriver.Open(connString))

	if err != nil {
		//dp.Logger.Error("%s", "GetPostgresConnection -> Open error: ", err.Error())
		return nil
	}

	log.Println("Postgres connection success: ", connString)
	return &postgres{Postgres: conn}
}

func (p *postgres) GetPostgresConnection() *gorm.DB {
	return p.Postgres
}
