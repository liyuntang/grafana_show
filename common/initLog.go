package common

import (
	"fmt"
	"io"
	"log"
	"os"
)

func InitLog(logFile string) *log.Logger {
	file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open logFile", logFile, "is bad err is", err)
		os.Exit(1)
	}
	return log.New(io.MultiWriter(file, os.Stdout), "[grafana_show] ", log.Ldate|log.Ltime|log.Lshortfile)
}
