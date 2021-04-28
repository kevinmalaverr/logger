package logger

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var colors = map[string]string{
	"success": "\x1b[32m",
	"info":    "\x1b[34m",
	"warn":    "\x1b[33m",
	"error":   "\x1b[31m",
	"default": "\x1b[0m",
}

var LogNames = map[string]string{
	"success": "SUCCESS",
	"info":    "INFO",
	"warn":    "WARN",
	"error":   "ERROR",
}

var shouldWrite = false

var logPath = ".log/"

func InitWriteLog() {
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		os.Mkdir(logPath, 0777)
	}

	shouldWrite = true
}

func WriteLog(message string) {
	logPath := logPath
	time := time.Now().Format("2006-01-02 15:04:05")

	f, err := os.OpenFile(logPath+time+".log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	if _, err := f.WriteString(time + " " + message + "\n"); err != nil {
		fmt.Println(err)
	}
}

func Log(logType string, message string) {
	fmt.Println(colors[logType], LogNames[logType], colors["default"], message)
	if shouldWrite {
		WriteLog(tabAlign(LogNames[logType], 0) + message)
	}
}

func Success(message string) {
	Log("success", message)
}

func Warn(message string) {
	Log("warn", message)
}

func Info(message string) {
	Log("info", message)
}

func Error(message string) {
	Log("error", message)
}

func tabAlign(str string, length int) string {
	if length == 0 {
		length = 10
	}
	return str + strings.Repeat(" ", length-len(str))
}
