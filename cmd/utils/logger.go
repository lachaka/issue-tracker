package utils

import (
	"log"
	"os"
)

type AppLogger struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

var instance *AppLogger
var Logger = *createLogger()

func createLogger() *AppLogger {

	if instance == nil {
		instance = &AppLogger {
			InfoLog : log.New(os.Stdout, "INFO\t", log.Ldate | log.Ltime),
			ErrorLog : log.New(os.Stderr, "ERROR\t", log.Ldate | log.Ltime | log.Lshortfile),
		}
	}

	return instance
} 