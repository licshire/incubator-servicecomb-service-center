// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build go1.9

package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const globalCallerSkip = 2

var (
	logger = NewLogger(Configure().WithCallerSkip(globalCallerSkip))

	zapLevelMap = map[string]zapcore.Level{
		"DEBUG": zap.DebugLevel,
		"INFO":  zap.InfoLevel,
		"WARN":  zap.WarnLevel,
		"ERROR": zap.ErrorLevel,
		"FATAL": zap.FatalLevel,
	}
)

func SetGlobal(cfg Config) {
	logger = NewLogger(cfg.WithCallerSkip(globalCallerSkip))

	// zap internal log
	_ = zap.ReplaceGlobals(logger.zapLogger)
	// golang log
	_ = zap.RedirectStdLog(logger.zapLogger)
}
