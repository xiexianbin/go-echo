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

// Package model implements the app GORM object by gorm.io/gorm.
package model

import (
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gorml "gorm.io/gorm/logger"

	"github.com/xiexianbin/go-echo-demo/config"
	"github.com/xiexianbin/go-echo-demo/pkg/log"
	"github.com/xiexianbin/go-echo-demo/pkg/util"
)

var (
	once sync.Once
	DB   *gorm.DB
)

func Init() {
	once.Do(func() {
		mysql_dsn := config.MSYQL_DSN
		if mysql_dsn != "" {
			conf := &gorm.Config{
				Logger: &GormLogger{
					ZapLogger: log.Logger,
					Config: gorml.Config{
						SlowThreshold:             200 * time.Microsecond,
						IgnoreRecordNotFoundError: true,
						LogLevel:                  gorml.Silent,
					},
				},
			}
			DB, err := gorm.Open(mysql.Open(mysql_dsn), conf)
			util.Mustf(err, "init MYSQL failed")

			// Connection Pool
			// ref https://gorm.io/docs/connecting_to_the_database.html#Connection-Pool
			sqlDB, err := DB.DB()
			util.Must(err)

			// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
			sqlDB.SetMaxIdleConns(10)

			// SetMaxOpenConns sets the maximum number of open connections to the database.
			sqlDB.SetMaxOpenConns(100)

			// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
			sqlDB.SetConnMaxLifetime(time.Hour)

			// Auto Migrate
			DB.AutoMigrate(
			// &User{},
			)
		}
	})
}

// Pool return *gorm.DB from Connection Pool
func Pool() *gorm.DB {
	return DB
}
