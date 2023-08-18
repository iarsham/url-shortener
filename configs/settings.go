package configs

import (
	"log"
	"os"
)

type CustomLogger struct {
	*log.Logger
}

func NewLogger() *CustomLogger {
	return &CustomLogger{
		Logger: log.New(os.Stdout, "[APP INFO]: ", log.Ldate|log.Ltime),
	}
}
