package logger

import (
	"fmt"
	"log"
	"os"
)


type loggerConfig struct {
	level     		string
	file      		*os.File
	stdLogger 		*log.Logger
	instanceID 		string
	projectRoot 	string  
	mode 			string
}

// 

func New(filePath string, level string, projectRoot string, mode string) Logger {
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	instanceID := fmt.Sprintf("%d", os.Getpid()) 

	if err != nil {
		panic(fmt.Sprintf("failed to open log file: %v", err))
	}

	return &loggerConfig{
		level:     level,
		file:      f,
		stdLogger: log.New(f, "", 0), // мы сами форматируем
		instanceID: instanceID,
		projectRoot: projectRoot,
		mode: 		 mode,
	}
}



