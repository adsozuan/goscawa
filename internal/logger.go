package goscawa

import (
	"io"
	"log"
	"os"
)

type AppLogger struct {
	*log.Logger
}

func NewAppLogger() *AppLogger {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	outputs := io.MultiWriter(file, os.Stdout)
	l := &AppLogger{}
	l.Logger = log.New(outputs, "INFO:", log.Ldate|log.Ltime|log.Lshortfile)
	return l
}

func (l *AppLogger) Info(message string) {
	l.SetPrefix("INFO:")
	l.Println(message)
}

func (l *AppLogger) Error(err error) {
	l.SetPrefix("ERROR:")
	l.Println(err)
}

func (l *AppLogger) Debug(message string) {
	l.SetPrefix("DEBUG:")
	l.Println(message)
}
