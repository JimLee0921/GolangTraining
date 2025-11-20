package main

import (
	"log"
	"os"
)

// 给系统定义多个 Logger

var (
	InfoLogger  *log.Logger
	WarnLogger  *log.Logger
	ErrorLogger *log.Logger
)

func main() {
	file, err := os.OpenFile("temp_files/app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	InfoLogger = log.New(file, "[INFO]  ", log.Ldate|log.Ltime|log.Lshortfile)
	WarnLogger = log.New(file, "[WARN]  ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)

	InfoLogger.Println("Server started")
	WarnLogger.Println("Memory usage is high")
	ErrorLogger.Println("Database connection failed")
}
