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

package log

import (
	"context"
	"time"

	"go.uber.org/zap"
	gorml "gorm.io/gorm/logger"
)

// GormLogger implemente gorm.io/gorm/logger/#Interface logging interface
type GormLogger struct {
	ZapLogger *zap.Logger
}

var _ gorml.Interface = (*GormLogger)(nil)

func (l *GormLogger) LogMode(level gorml.LogLevel) gorml.Interface {
	// TODO(xiexianbin): set log level
	return l
}
func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.ZapLogger.Sugar().Infof(msg, data...)
}

func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.ZapLogger.Sugar().Warnf(msg, data...)
}

func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.ZapLogger.Sugar().Errorf(msg, data...)
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil {
		l.ZapLogger.Sugar().Errorf("%s [%.3fms] %s rows:%v error:%v", "query", float64(elapsed.Nanoseconds())/1e6, sql, rows, err)

	} else {
		l.ZapLogger.Sugar().Infof("%s [%.3fms] %s rows:%v", "query", float64(elapsed.Nanoseconds())/1e6, sql, rows)
	}
}
