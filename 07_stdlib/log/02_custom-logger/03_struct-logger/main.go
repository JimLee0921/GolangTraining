package main

import (
	"io"
	"log"
	"os"
)

// MyLogger 封装为结构体，便于外部调用
type MyLogger struct {
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
}

func NewMyLogger(w io.Writer) *MyLogger {
	return &MyLogger{
		Info:  log.New(w, "[INFO]  ", log.Ldate|log.Ltime|log.Lshortfile),
		Warn:  log.New(w, "[WARN]  ", log.Ldate|log.Ltime|log.Lshortfile),
		Error: log.New(w, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func main() {
	file, _ := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	logger := NewMyLogger(file)

	logger.Info.Println("Server started")
	logger.Warn.Println("High memory usage")
	logger.Error.Println("Database failed")
}
