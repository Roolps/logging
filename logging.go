package logging

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

var (
	debug = false

	white  = color.New(color.FgWhite)
	grey   = color.New(color.FgWhite, color.Faint)
	yellow = color.New(color.FgYellow)
	red    = color.New(color.FgRed)
)

func EnableDebug() {
	debug = true
}

func Debugf(format string, args ...any) {
	if debug {
		grey.Println(time.Now().Format("2006/01/02 15:04:05 [DEBUG] "), fmt.Sprintf(format, args...))
	}
}

func Debug(content any) {
	if debug {
		grey.Println(time.Now().Format("2006/01/02 15:04:05 [DEBUG] "), content)
	}
}

func Infof(format string, args ...any) {
	white.Println(time.Now().Format("2006/01/02 15:04:05 [INFO] "), fmt.Sprintf(format, args...))
}

func Info(content any) {
	white.Println(time.Now().Format("2006/01/02 15:04:05 [INFO] "), content)
}

func Warnf(format string, args ...any) {
	yellow.Println(time.Now().Format("2006/01/02 15:04:05 [WARN] "), fmt.Sprintf(format, args...))
}

func Warn(content any) {
	yellow.Println(time.Now().Format("2006/01/02 15:04:05 [WARN] "), content)
}

func Errorf(format string, args ...any) {
	red.Println(time.Now().Format("2006/01/02 15:04:05 [ERROR] "), fmt.Sprintf(format, args...))
}

func Error(content any) {
	red.Println(time.Now().Format("2006/01/02 15:04:05 [ERROR] "), content)
}
