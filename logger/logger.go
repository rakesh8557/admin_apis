package logger

import (
	"admin_apis/config"
	"log"

	"gopkg.in/natefinch/lumberjack.v2"
)

var logFile *lumberjack.Logger

var logger *log.Logger

func InitLogger() {
	logFile = &lumberjack.Logger{
		Filename:   *config.LOG_FILE_PATH,
		MaxSize:    5, // MB
		MaxBackups: 3,
		MaxAge:     28, // days
		Compress:   true,
	}

	logger = log.New(logFile, "", log.LstdFlags|log.Lshortfile)
}

func Log(level string, message string) {
	prefix := "[" + level + "] "
	logger.Println(prefix + message)
}
