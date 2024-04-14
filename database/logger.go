package database

import "log"

func initLogger() *log.Logger {
	logger := log.Logger{}
	logger.SetPrefix("[SQL]")
	return &logger
}
