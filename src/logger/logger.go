package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

const (
	clr_0 = "\x1b[30m"
	clr_R = "\x1b[31m"
	clr_G = "\x1b[32m"
	clr_Y = "\x1b[33m"
	clr_B = "\x1b[34m"
	clr_M = "\x1b[35m"
	clr_C = "\x1b[36m"
	clr_W = "\x1b[37m"
	clr_N = "\x1b[0m"
)

type Logger struct {
	logger *log.Logger
}

type LogLevel int

const layout = "2006-01-02T15:04:05.000Z"

const (
	info LogLevel = 1 + iota
	debug
	err
)

var localLog Logger
var (
	logLvls = []string{"INFO", "DEBUG", "ERROR"}
)

func (ll LogLevel) String() string { return logLvls[ll-1] }

/*
NewLogger creates a logger instance based on the name provided
*/
func NewLogger(name string) *Logger {
	localLog = Logger{logger: log.New(os.Stderr, `[`+name+`]`, 0)}
	return &localLog
}

/*
GetLogger returns already created logger instance
*/
func GetLogger() *Logger {
	return &localLog
}

/*
Error logs the error with the arguments provided
*/
func (l *Logger) Error(values ...interface{}) {
	l.out(err, "", values...)
}

/*
Errorf logs the error with the formating string
*/
func (l *Logger) Errorf(format string, values ...interface{}) {
	l.out(err, format, values...)
}

/*
Info logs the information with the arguments provided
*/
func (l *Logger) Info(values ...interface{}) {
	l.out(info, "", values...)
}

/*
Infof logs the information with the formating string
*/
func (l *Logger) Infof(format string, values ...interface{}) {
	l.out(info, format, values...)
}

/*
Debug logs the debug information with the arguments provided
*/
func (l *Logger) Debug(values ...interface{}) {
	l.out(debug, "", values...)
}

/*
Debugf logs the information with the formating string
*/
func (l *Logger) Debugf(format string, values ...interface{}) {
	l.out(debug, format, values...)
}

func (l *Logger) out(level LogLevel, format string, values ...interface{}) {
	var str string
	var color string
	if format == "" {
		str = fmt.Sprint(values...)
	} else {
		str = fmt.Sprintf(format, values...)
	}
	switch level {
	case info:
		color = clr_G
	case debug:
		color = clr_B
	case err:
		color = clr_R
	}
	fmt.Printf("%s", color)
	pc, fn, line, _ := runtime.Caller(2)
	if runtime.FuncForPC(pc).Name() == "main.Error" || runtime.FuncForPC(pc).Name() == "main.Errorf" {
		pc, fn, line, _ = runtime.Caller(3)
	}
	//outStr := fmt.Sprintf(`,"LogLevel":"%s","Timestamp":"%s","Action":"[%s][%s:%d]","Message":"%s"}`, level.String(), time.Now().Format(layout), runtime.FuncForPC(pc).Name(), fn, line, str)
	outStr := fmt.Sprintf(` [%s] %s [%s][%s:%d] %s -> `, level.String(), time.Now().Format(layout), runtime.FuncForPC(pc).Name(), fn, line, str)
	l.logger.Printf(outStr)
	fmt.Printf("%s", clr_N)
}
