package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

const (
	MYSQL_HOST        = "MYSQL_HOST"
	MYSQL_PORT        = "MYSQL_PORT"
	MYSQL_USER        = "MYSQL_USER"
	MYSQL_PASS        = "MYSQL_PASSWORD"
	MYSQL_DATABASE    = "MYSQL_DATABASE"
	MYSQL_LOC         = "MYSQL_LOC"
	DEFAULT_COLLATION = "utf8mb4_unicode_ci"
	LOC_JP            = "Asia/Tokyo"
)

func Get() mysql.Config {
	conf := mysql.Config{}
	hostname := getHostname()
	port := getPort()
	conf.Addr = fmt.Sprintf("%s:%s", hostname, port)
	setUser(&conf)
	setPasswd(&conf)
	setDatabase(&conf)
	conf.Loc = getLocation()
	conf.Net = "tcp"
	return conf
}

func getPort() string {
	var port string
	if p, ok := os.LookupEnv(MYSQL_PORT); ok {
		port = p
	} else {
		port = "3306"
	}
	return port
}

func getHostname() string {
	var hostname string
	if h, ok := os.LookupEnv(MYSQL_HOST); ok {
		hostname = h
	} else {
		hostname = "database"
	}
	return hostname
}

func setUser(conf *mysql.Config) {
	if username, ok := os.LookupEnv(MYSQL_USER); ok {
		conf.User = username
	} else {
		keyfatal(MYSQL_USER)
	}
}
func setPasswd(conf *mysql.Config) {
	if password, ok := os.LookupEnv(MYSQL_PASS); ok {
		conf.Passwd = password
	} else {
		keyfatal(MYSQL_PASS)
	}
}

func setDatabase(conf *mysql.Config) {
	if database, ok := os.LookupEnv(MYSQL_DATABASE); ok {
		conf.DBName = database
	} else {
		keyfatal(MYSQL_DATABASE)
	}
}

func getLocation() *time.Location {
	var (
		loc *time.Location
		err error
	)
	if l, ok := os.LookupEnv(MYSQL_LOC); ok {
		loc, err = time.LoadLocation(l)
	} else {
		loc, err = time.LoadLocation(LOC_JP)
	}
	if err != nil {
		log.Fatalln(err)
	}
	return loc
}
func keyfatal(key string) {
	log.Fatalf("please set env: %s\n", key)
}
