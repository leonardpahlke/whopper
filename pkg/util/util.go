// Copyright 2022 Leonard Vincent Simon Pahlke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// GetLogger this function is used to get a logger which is used to produce log outputs
func GetLogger(lvl zapcore.Level) *zap.SugaredLogger {
	rawJSON := []byte(`{
		"level": "` + lvl.String() + `",
		"encoding": "json",
		"outputPaths": ["stdout"],
		"errorOutputPaths": ["stderr"],
		"encoderConfig": {"messageKey": "message","levelKey": "level","levelEncoder": "lowercase"}
	  }`,
	)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(fmt.Sprintf("[util/GetLogger] error in json.Unmarshal %v", err))
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(fmt.Sprintf("error in cfg.Build %v", err))
	}
	sugar := logger.Sugar()
	return sugar
}

// MatchLogLevel used to get the zap log level
func MatchLogLevel(lvl WrapLogLevel) zapcore.Level {
	switch lvl {
	case Debug:
		return zapcore.DebugLevel
	case Info:
		return zapcore.InfoLevel
	case Warning:
		return zapcore.WarnLevel
	case Error:
		return zapcore.ErrorLevel
	case Dpanic:
		return zapcore.DPanicLevel
	case Panic:
		return zapcore.PanicLevel
	case Fatal:
		return zapcore.FatalLevel
	}
	return zapcore.DebugLevel
}

// WrapLogLevel defines zap log levels
type WrapLogLevel string

const (
	// Debug log level (-1)
	Debug WrapLogLevel = "DEBUG"
	// Info log level (0)
	Info WrapLogLevel = "INFO"
	// Warning log level (1)
	Warning WrapLogLevel = "WARNING"
	// Error log level (2)
	Error WrapLogLevel = "ERROR"
	// Dpanic log level (3)
	Dpanic WrapLogLevel = "DPANIC"
	// Panic log level (4)
	Panic WrapLogLevel = "PANIC"
	// Fatal log level /5
	Fatal WrapLogLevel = "FATAL"
)
