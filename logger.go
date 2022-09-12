/*
 * Copyright 2022 Yann Dumont
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package log_level

import (
	"errors"
	"fmt"
	"github.com/y-du/go-log-level/level"
	"log"
)

type Logger struct {
	level  level.Level
	pLevel level.Level
	*log.Logger
}

func New(logger *log.Logger, level level.Level) (*Logger, error) {
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

func (l *Logger) GetLevel() level.Level {
	return l.level
}

func (l *Logger) SetPrintLevel(level level.Level) (err error) {
	if err = checkLevel(level); err != nil {
		return err
	}
	l.pLevel = level
	return
}

func (l *Logger) output(level level.Level, v string) (err error) {
	if level <= l.level {
		return l.Output(3, v)
	}
	return
}

func checkLevel(l level.Level) error {
	if l >= level.Off && l <= level.Debug {
		return nil
	}
	return errors.New(fmt.Sprintf("unknown logging level '%d'", l))
}
