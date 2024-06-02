// Copyright 2024 xiexianbin<me@xiexianbin.cn>
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

// ref https://github.com/argoproj/argo-workflows/blob/v3.5.7/util/env/env.go
package util

import (
	"os"
	"strconv"
	"time"

	"go.uber.org/zap"

	"github.com/xiexianbin/go-echo-demo/pkg/log"
)

func LookupEnvDurationOr(key string, o time.Duration) time.Duration {
	v, found := os.LookupEnv(key)
	if found && v != "" {
		d, err := time.ParseDuration(v)
		if err != nil {
			log.CoreLogger.Sugar().Panicw("failed to parse", zap.Any(key, v), zap.Error(err))
		} else {
			return d
		}
	}
	return o
}

func LookupEnvIntOr(key string, o int) int {
	v, found := os.LookupEnv(key)
	if found && v != "" {
		d, err := strconv.Atoi(v)
		if err != nil {
			log.CoreLogger.Sugar().Panicw("failed to convert to int", zap.Any(key, v), zap.Error(err))
		} else {
			return d
		}
	}
	return o
}

func LookupEnvFloatOr(key string, o float64) float64 {
	v, found := os.LookupEnv(key)
	if found && v != "" {
		d, err := strconv.ParseFloat(v, 64)
		if err != nil {
			log.CoreLogger.Sugar().Panicw("failed to convert to float", zap.Any(key, v), zap.Error(err))
		} else {
			return d
		}
	}
	return o
}

func LookupEnvStringOr(key string, o string) string {
	v, found := os.LookupEnv(key)
	if found && v != "" {
		return v
	}
	return o
}
