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

package level

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Level int8

const (
	Off Level = iota
	Error
	Warning
	Info
	Debug
	Default = Warning
)

var levelStr = [5]string{
	"off",
	"error",
	"warning",
	"info",
	"debug",
}

func (l Level) String() string {
	return levelStr[l]
}

func (l *Level) UnmarshalJSON(data []byte) (err error) {
	var itf interface{}
	if err = json.Unmarshal(data, &itf); err != nil {
		return
	}
	*l, err = Parse(itf.(string))
	return
}

func Parse(v string) (Level, error) {
	for i := 0; i < len(levelStr); i++ {
		if levelStr[i] == v {
			return Level(i), nil
		}
	}
	return Default, errors.New(fmt.Sprintf("unknown logging level '%s'", v))
}
