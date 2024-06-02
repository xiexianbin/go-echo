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
	"io"

	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
)

// EchoLogger implemente echo.Logger logging interface
type EchoLogger struct {
	ZapLogger *zap.Logger
}

func (l *EchoLogger) Output() io.Writer {
	return nil
}

func (l *EchoLogger) SetOutput(w io.Writer) {
}

func (l *EchoLogger) Prefix() string {
	return ""
}

func (l *EchoLogger) SetPrefix(p string) {
	l.Debugf("un-support set prefix: %s", p)
}

func (l *EchoLogger) Level() log.Lvl {
	return log.Lvl(l.ZapLogger.Sugar().Level())
}

func (l *EchoLogger) SetLevel(v log.Lvl) {
}

func (l *EchoLogger) SetHeader(h string) {
}

func (l *EchoLogger) Print(i ...interface{}) {
	l.ZapLogger.Sugar().Debug(i[0].(string))
}

func (l *EchoLogger) Printf(format string, args ...interface{}) {
	l.ZapLogger.Sugar().Debugf(format, args...)
}

func (l *EchoLogger) Printj(j log.JSON) {
	l.Debugj(j)
}

func (l *EchoLogger) Debug(i ...interface{}) {
	l.ZapLogger.Sugar().Debug(i)
}

func (l *EchoLogger) Debugf(format string, args ...interface{}) {
	l.ZapLogger.Sugar().Debugf(format, args...)
}

func (l *EchoLogger) Debugj(j log.JSON) {
	l.ZapLogger.Sugar().Debugw("", jsonToZapFields(j))
}

func (l *EchoLogger) Info(i ...interface{}) {
	l.ZapLogger.Sugar().Info(i)
}

func (l *EchoLogger) Infof(format string, args ...interface{}) {
	l.ZapLogger.Sugar().Infof(format, args...)
}

func (l *EchoLogger) Infoj(j log.JSON) {
	l.ZapLogger.Sugar().Infow("", jsonToZapFields(j))
}

func (l *EchoLogger) Warn(i ...interface{}) {
	l.ZapLogger.Sugar().Warn(i)
}

func (l *EchoLogger) Warnf(format string, args ...interface{}) {
	l.ZapLogger.Sugar().Warnf(format, args...)
}

func (l *EchoLogger) Warnj(j log.JSON) {
	l.ZapLogger.Sugar().Warnw("", jsonToZapFields(j))
}

func (l *EchoLogger) Error(i ...interface{}) {
	l.ZapLogger.Sugar().Error(i)
}

func (l *EchoLogger) Errorf(format string, args ...interface{}) {
	l.ZapLogger.Sugar().Errorf(format, args...)
}

func (l *EchoLogger) Errorj(j log.JSON) {
	l.ZapLogger.Sugar().Errorw("", jsonToZapFields(j))
}

func (l *EchoLogger) Fatal(i ...interface{}) {
	l.ZapLogger.Sugar().Fatal(i)
}

func (l *EchoLogger) Fatalj(j log.JSON) {
	l.ZapLogger.Sugar().Fatalw("", jsonToZapFields(j))
}

func (l *EchoLogger) Fatalf(format string, args ...interface{}) {
	l.ZapLogger.Sugar().Fatalf(format, args...)
}

func (l *EchoLogger) Panic(i ...interface{}) {
	l.ZapLogger.Sugar().Panic(i)
}

func (l *EchoLogger) Panicj(j log.JSON) {
	l.ZapLogger.Sugar().Panicw("", jsonToZapFields(j))
}

func (l *EchoLogger) Panicf(format string, args ...interface{}) {
	l.ZapLogger.Sugar().Panicf(format, args...)
}

func jsonToZapFields(j log.JSON) []zap.Field {
	fields := make([]zap.Field, 0)
	for k, v := range j {
		switch v.(type) {
		case string:
			fields = append(fields, zap.String(k, v.(string)))
		default:
			fields = append(fields, zap.Any(k, v))
		}
	}
	return fields
}
