package log_level

import (
	"errors"
	"fmt"
	"github.com/y-du/go-log-level/level"
	"log"
)

type Logger struct {
	level  int
	pLevel int
	*log.Logger
}

func New(logger *log.Logger, level int) (*Logger, error) {
	if err := checkLevel(level); err != nil {
		return nil, err
	}
	return &Logger{
		level:  level,
		pLevel: -1,
		Logger: logger,
	}, nil
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.output(l.pLevel, fmt.Sprintf(format, v...))
}

func (l *Logger) Print(v ...interface{}) {
	l.output(l.pLevel, fmt.Sprint(v...))
}

func (l *Logger) Println(v ...interface{}) {
	l.output(l.pLevel, fmt.Sprintln(v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.output(level.Error, fmt.Sprint(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.output(level.Error, fmt.Sprintf(format, v...))
}

func (l *Logger) Errorln(v ...interface{}) {
	l.output(level.Error, fmt.Sprintln(v...))
}

func (l *Logger) Warning(v ...interface{}) {
	l.output(level.Warning, fmt.Sprint(v...))
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	l.output(level.Warning, fmt.Sprintf(format, v...))
}

func (l *Logger) Warningln(v ...interface{}) {
	l.output(level.Warning, fmt.Sprintln(v...))
}

func (l *Logger) Info(v ...interface{}) {
	l.output(level.Info, fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.output(level.Info, fmt.Sprintf(format, v...))
}

func (l *Logger) Infoln(v ...interface{}) {
	l.output(level.Info, fmt.Sprintln(v...))
}

func (l *Logger) Debug(v ...interface{}) {
	l.output(level.Debug, fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.output(level.Debug, fmt.Sprintf(format, v...))
}

func (l *Logger) Debugln(v ...interface{}) {
	l.output(level.Debug, fmt.Sprintln(v...))
}

func (l *Logger) GetLevel() int {
	return l.level
}

func (l *Logger) SetPrintLevel(level int) (err error) {
	if err = checkLevel(level); err != nil {
		return err
	}
	l.pLevel = level
	return
}

func (l *Logger) output(level int, v string) (err error) {
	if level <= l.level {
		return l.Output(3, v)
	}
	return
}

var lvlStrings = [5]string{
	"off",
	"error",
	"warning",
	"info",
	"debug",
}

func ParseLevel(v string) (int, error) {
	for i := 0; i < len(lvlStrings); i++ {
		if lvlStrings[i] == v {
			return i, nil
		}
	}
	return level.Default, errors.New(fmt.Sprintf("unknown logging level '%s'", v))
}

func checkLevel(l int) error {
	if l >= level.Off && l <= level.Debug {
		return nil
	}
	return errors.New(fmt.Sprintf("unknown logging level '%d'", l))
}
