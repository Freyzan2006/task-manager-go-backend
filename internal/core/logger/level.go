package logger 

import (
	"os"
)

// Реализация интерфейса
func (l *loggerConfig) Debug(v ...interface{}) { l.log("DEBUG", colorGreen, v...) }
func (l *loggerConfig) Info(v ...interface{})  { l.log("INFO", colorBlue, v...) }
func (l *loggerConfig) Warn(v ...interface{})  { l.log("WARN", colorYellow, v...) }
func (l *loggerConfig) Error(v ...interface{}) { l.log("ERROR", colorRed, v...) }
func (l *loggerConfig) Fatal(v ...interface{}) {
	l.log("FATAL", colorRed, v...)
	os.Exit(1)
}


func (l *loggerConfig) Infof(format string, v ...interface{})  { l.logf("INFO", colorBlue, format, v...) }
func (l *loggerConfig) Debugf(format string, v ...interface{}) { l.logf("DEBUG", colorGreen, format, v...) }
func (l *loggerConfig) Warnf(format string, v ...interface{})  { l.logf("WARN", colorYellow, format, v...) }
func (l *loggerConfig) Errorf(format string, v ...interface{}) { l.logf("ERROR", colorRed, format, v...) }
func (l *loggerConfig) Fatalf(format string, v ...interface{}) {
	l.logf("FATAL", colorRed, format, v...)
	os.Exit(1)
}