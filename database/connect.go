package database

import (
	"database/sql"
	"log"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/kynmh69/go-mysql/config"
)

const DRIVER = "mysql"

var Connection *goqu.Database

func ConnectToMySQL() {
	conf := config.Get()
	db, err := sql.Open(DRIVER, conf.FormatDSN())
	if err != nil {
		log.Fatalln("cannot open db", err)
	}
	if err = db.Ping(); err != nil {
		db.Close()
		log.Fatalln("cannot connect db.")
	}
	Connection = goqu.New(DRIVER, db)
}
