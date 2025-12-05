package mylog

import (
	"io"
	"log"
	"os"
	"sync"
)

/*
创建两个日志实例分别用于打印 Info 和 Error 日志
*/

var (
	// 第一个参数是 日志写到哪里，第二个参数是日志每一行前缀字符串，第三个参数是控制日志前缀格式（时间、日期、文件名等）
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorLog, infoLog}
	mu       sync.Mutex
)

// log 方法
var (
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Info   = infoLog.Print
	Infof  = infoLog.Printf
)

// 日志层级设计
const (
	InfoLevel = iota
	ErrorLevel
	Disabled
)

// SetLevel 控制日志层级，level 决定哪些级别的日志需要被舍弃
func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if ErrorLevel < level {
		errorLog.SetOutput(io.Discard)
	}

	if InfoLevel < level {
		infoLog.SetOutput(io.Discard)
	}
}
