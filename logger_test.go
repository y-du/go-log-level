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
	"bytes"
	"fmt"
	"github.com/y-du/go-log-level/level"
	"log"
	"strings"
	"testing"
)

func testParseLevel(t *testing.T, a level.Level) {
	if b, err := level.Parse(a.String()); err != nil || b != a {
		t.Errorf("b = %d; want %d", b, a)
	}
}

func TestParseLevels(t *testing.T) {
	testParseLevel(t, level.Off)
	testParseLevel(t, level.Error)
	testParseLevel(t, level.Warning)
	testParseLevel(t, level.Info)
	testParseLevel(t, level.Debug)
	testParseLevel(t, level.Default)
}

func TestGetLevel(t *testing.T) {
	l, _ := New(log.Default(), level.Error)
	b := l.GetLevel()
	if b != level.Error {
		t.Errorf("b = %d; want %d", b, level.Error)
	}
}

func testOutput(t *testing.T, lvl level.Level, a string) {
	buf := new(bytes.Buffer)
	l, _ := New(log.New(buf, "", 0), lvl)
	l.Error(level.Error)
	l.Warning(level.Warning)
	l.Info(level.Info)
	l.Debug(level.Debug)
	b := strings.ReplaceAll(buf.String(), "\n", "")
	if b != a {
		t.Errorf("b = %s; want %s", b, a)
	}
}

func TestOff(t *testing.T) {
	testOutput(t, level.Off, "")
}

func TestError(t *testing.T) {
	testOutput(t, level.Error, level.Error.String())
}

func TestWarning(t *testing.T) {
	testOutput(t, level.Warning, fmt.Sprint(level.Error.String(), level.Warning.String()))
}

func TestInfo(t *testing.T) {
	testOutput(t, level.Info, fmt.Sprint(level.Error.String(), level.Warning.String(), level.Info.String()))
}

func TestDebug(t *testing.T) {
	testOutput(t, level.Debug, fmt.Sprint(level.Error.String(), level.Warning.String(), level.Info.String(), level.Debug.String()))
}

func TestSetPrintLevel(t *testing.T) {
	buf := new(bytes.Buffer)
	l, _ := New(log.New(buf, "", 0), level.Error)
	a := "test"
	l.Print(a)
	b := strings.ReplaceAll(buf.String(), "\n", "")
	if b != a {
		t.Errorf("b = %s; want %s", b, a)
	}
	buf.Reset()
	l.SetPrintLevel(level.Info)
	l.Print(a)
	b = strings.ReplaceAll(buf.String(), "\n", "")
	if b == a {
		t.Errorf("b = %s; want %s", b, "")
	}
}
