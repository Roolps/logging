package logging

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/fatih/color"
)

var (
	debug = false

	white  = color.New(color.FgWhite)
	grey   = color.New(color.FgWhite, color.Faint)
	yellow = color.New(color.FgYellow)
	red    = color.New(color.FgRed)
	cyan   = color.New(color.FgCyan)

	errorChnl chan string
	logChnl   chan string

	prefix string
)

type Profile struct {
	Prefix string
	Color  *color.Color
	debug  bool
}

func (p *Profile) EnableDebug() {
	p.debug = true
}

func SetPrefix(p string, c *color.Color) {
	prefix = c.Sprintf("[%v]", p)
}

func openFile(f string) (*os.File, error) {
	_, err := os.Stat(f)

	if os.IsNotExist(err) {
		dir := filepath.Dir(f)

		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return nil, fmt.Errorf("failed to create directories: %w", err)
		}

		file, err := os.Create(f)
		if err != nil {
			return nil, fmt.Errorf("failed to create file: %w", err)
		}
		file.Close()
	}

	file, err := os.OpenFile(f, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open file in append mode: %w", err)
	}

	return file, nil
}

// TO DO : make this for different types :)
func fileListner(file *os.File, chnl chan string) {
	defer file.Close()
	for {
		msg, ok := <-chnl
		if !ok {
			return
		}
		if _, err := file.WriteString(msg + "\n"); err != nil {
			fmt.Println(
				grey.Sprint(time.Now().Format("02.01.2006 15:04:05")),
				white.Sprintf("[Logging]"),
				red.Sprint("ERROR"),
				white.Sprint(err),
			)
		}
	}
}

func SetErrorFile(f string) error {
	errorChnl = make(chan string)

	file, err := openFile(f)
	if err != nil {
		return err
	}
	go fileListner(file, errorChnl)
	return nil
}

func SetLogFile(f string) error {
	logChnl = make(chan string)

	file, err := openFile(f)
	if err != nil {
		return err
	}
	go fileListner(file, logChnl)
	return nil
}
func (p *Profile) Debug(content any) {
	if p.debug {
		fmt.Println(
			grey.Sprint(time.Now().Format("02.01.2006 15:04:05")),
			p.Color.Sprintf("[%v]", p.Prefix),
			white.Sprint("DEBUG"),
			grey.Sprint(content),
		)
		if logChnl != nil {
			logChnl <- fmt.Sprintf("%s %s DEBUG %v", time.Now().Format("02.01.2006 15:04:05"), p.Prefix, content)
		}
	}
}

func (p *Profile) Debugf(format string, args ...any) {
	if p.debug {
		fmt.Println(
			grey.Sprint(time.Now().Format("02.01.2006 15:04:05")),
			p.Color.Sprintf("[%v]", p.Prefix),
			white.Sprint("DEBUG"),
			grey.Sprint(fmt.Sprintf(format, args...)),
		)
		if logChnl != nil {
			logChnl <- fmt.Sprintf("%s %s DEBUG %s", time.Now().Format("02.01.2006 15:04:05"), p.Prefix, fmt.Sprintf(format, args...))
		}
	}
}

func (p *Profile) Info(content any) {
	fmt.Println(
		grey.Sprint(time.Now().Format("02.01.2006 15:04:05")),
		p.Color.Sprintf("[%v]", p.Prefix),
		cyan.Sprint("INFO"),
		white.Sprint(content),
	)
	if logChnl != nil {
		logChnl <- fmt.Sprintf("%s %s INFO %v", time.Now().Format("02.01.2006 15:04:05"), p.Prefix, content)
	}


func (p *Profile) Infof(format string, args ...any) {
	fmt.Println(
		grey.Sprint(time.Now().Format("02.01.2006 15:04:05")),
		p.Color.Sprintf("[%v]", p.Prefix),
		cyan.Sprint("INFO"),
		white.Sprint(fmt.Sprintf(format, args...)),
	)
	if logChnl != nil {
		logChnl <- fmt.Sprintf("%s %s INFO %s", time.Now().Format("02.01.2006 15:04:05"), p.Prefix, fmt.Sprintf(format, args...))
	}
}

func (p *Profile) Warn(content any) {
	fmt.Println(
		grey.Sprint(time.Now().Format("02.01.2006 15:04:05")),
		p.Color.Sprintf("[%v]", p.Prefix),
		yellow.Sprint("WARN"),
		white.Sprint(content),
	)
	if logChnl != nil {
		logChnl <- fmt.Sprintf("%s [%s] WARN %v", time.Now().Format("02.01.2006 15:04:05"), p.Prefix, content)
	}
}

func (p *Profile) Warnf(format string, args ...any) {
	content := fmt.Sprintf(format, args...)
	fmt.Println(
		grey.Sprint(time.Now().Format("02.01.2006 15:04:05")),
		p.Color.Sprintf("[%v]", p.Prefix),
		yellow.Sprint("WARN"),
		white.Sprintf(format, args...),
	)
	if logChnl != nil {
		logChnl <- fmt.Sprintf("%s [%s] WARN %v", time.Now().Format("02.01.2006 15:04:05"), p.Prefix, content)
	}
}

func (p *Profile) Error(content any) {
	_, file, line, _ := runtime.Caller(1)

	fmt.Println(
		grey.Sprint(time.Now().Format("02.01.2006 15:04:05")),
		p.Color.Sprintf("[%v]", p.Prefix),
		red.Sprintf("ERROR %v:%v", filepath.Base(file), line),
		white.Sprint(content),
	)
	if errorChnl != nil {
		errorChnl <- fmt.Sprintf("%s [%s] ERROR %v", time.Now().Format("02.01.2006 15:04:05"), p.Prefix, content)
	}
	if logChnl != nil {
		logChnl <- fmt.Sprintf("%s [%s] ERROR %v", time.Now().Format("02.01.2006 15:04:05"), p.Prefix, content)
	}
}

func (p *Profile) Errorf(format string, args ...any) {
	_, file, line, _ := runtime.Caller(1)

	content := fmt.Sprintf(format, args...)
	fmt.Println(
		grey.Sprint(time.Now().Format("02.01.2006 15:04:05")),
		p.Color.Sprintf("[%v]", p.Prefix),
		red.Sprintf("ERROR %v:%v", filepath.Base(file), line),
		white.Sprint(content),
	)
	if errorChnl != nil {
		errorChnl <- fmt.Sprintf("%s [%s] ERROR %v", time.Now().Format("02.01.2006 15:04:05"), p.Prefix, content)
	}
	if logChnl != nil {
		logChnl <- fmt.Sprintf("%s [%s] ERROR %v", time.Now().Format("02.01.2006 15:04:05"), p.Prefix, content)
	}
}
