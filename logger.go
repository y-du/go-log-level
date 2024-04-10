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
	"fmt"
	"github.com/y-du/go-log-level/level"
	"log"
)

type Logger struct {
	level   level.Level
	pLevel  level.Level
	ePrefix string
	wPrefix string
	iPrefix string
	dPrefix string
	*log.Logger
}

func New(logLogger *log.Logger, loggerLevel level.Level) (l *Logger, err error) {
	if ok := checkLevel(loggerLevel); !ok {
		err = fmt.Errorf("unknown level '%d': defaulting to '%s'", loggerLevel, level.Default)
		loggerLevel = level.Default
	}
	l = &Logger{
		level:  loggerLevel,
		pLevel: -1,
		Logger: logLogger,
	}
	return
}

func (l *Logger) SetLevelPrefix(error, warning, info, debug string) {
	l.ePrefix = error
	l.wPrefix = warning
	l.iPrefix = info
	l.dPrefix = debug
}

func (l *Logger) Printf(format string, v ...interface{}) {
	if l.pLevel <= l.level {
		l.Output(3, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Print(v ...interface{}) {
	if l.pLevel <= l.level {
		l.Output(3, fmt.Sprint(v...))
	}
}

func (l *Logger) Println(v ...interface{}) {
	if l.pLevel <= l.level {
		l.Output(3, fmt.Sprint(v...))
	}
}

func (l *Logger) Error(v ...interface{}) {
	if level.Error <= l.level {
		l.Output(3, l.ePrefix+fmt.Sprint(v...))
	}
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	if level.Error <= l.level {
		l.Output(3, l.ePrefix+fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Warning(v ...interface{}) {
	if level.Warning <= l.level {
		l.Output(3, l.wPrefix+fmt.Sprint(v...))
	}
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	if level.Warning <= l.level {
		l.Output(3, l.wPrefix+fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Info(v ...interface{}) {
	if level.Info <= l.level {
		l.Output(3, l.iPrefix+fmt.Sprint(v...))
	}
}

func (l *Logger) Infof(format string, v ...interface{}) {
	if level.Info <= l.level {
		l.Output(3, l.iPrefix+fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Debug(v ...interface{}) {
	if level.Debug <= l.level {
		l.Output(3, l.dPrefix+fmt.Sprint(v...))
	}
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if level.Debug <= l.level {
		l.Output(3, l.dPrefix+fmt.Sprintf(format, v...))
	}
}

func (l *Logger) GetLevel() level.Level {
	return l.level
}

func (l *Logger) SetPrintLevel(level level.Level) (err error) {
	if ok := checkLevel(level); !ok {
		return fmt.Errorf("unknown level '%d'", level)
	}
	l.pLevel = level
	return
}

func checkLevel(l level.Level) bool {
	if l >= level.Off && l <= level.Debug {
		return true
	}
	return false
}
