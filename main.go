package main

import (
	"log"

	"github.com/kynmh69/go-mysql/database"
)

func init() {
	database.ConnectToMySQL()
}

func main() {
	log.Println("start app")
}
