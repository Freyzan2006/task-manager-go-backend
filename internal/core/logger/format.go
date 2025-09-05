package logger 

import (
	"fmt"
	"runtime"
	"time"
)

import (
	"path/filepath"
	"strings"
)


// Цвета ANSI для консоли
var (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[36m"
	colorGreen  = "\033[32m"
)



// общий метод логирования
func (l *loggerConfig) log(level string, color string, v ...interface{}) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "unknown"
		line = 0
	} else {
		// делаем путь относительным к projectRoot
		if idx := strings.Index(file, l.projectRoot); idx != -1 {
			file = file[idx+len(l.projectRoot)+1:] // обрежем всё до "internal/..."
		} else {
			file = filepath.Base(file) // fallback: только имя файла
		}
	}

	msg := fmt.Sprint(v...)
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	formatted := fmt.Sprintf(
		"[PID:%s] [%s] [%-5s] [%s] %s (%s:%d)",
		l.instanceID, l.mode, level, timestamp, msg, file, line,
	)

	// Печать в консоль с цветом
	fmt.Println(color + formatted + colorReset)

	// Печать в файл без цвета
	l.stdLogger.Println(formatted)
}


func (l *loggerConfig) logf(level string, color string, format string, v ...interface{}) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "unknown"
		line = 0
	} else {
		// делаем путь относительным к projectRoot
		if idx := strings.Index(file, l.projectRoot); idx != -1 {
			file = file[idx+len(l.projectRoot)+1:] // обрежем всё до "internal/..."
		} else {
			file = filepath.Base(file) // fallback: только имя файла
		}
	}

	msg := fmt.Sprintf(format, v...)
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	formatted := fmt.Sprintf(
		"[PID:%s] [%s] [%-5s] [%s] %s (%s:%d)",
		l.instanceID, l.mode, level, timestamp, msg, file, line,
	)

	// Цветная печать в консоль
	fmt.Println(color + formatted + colorReset)

	// Запись в файл без цвета
	l.stdLogger.Println(formatted)
}
