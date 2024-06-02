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

package config

import (
	"github.com/xiexianbin/go-echo-demo/pkg/util"
)

var (
	// log level, support debug, info, warn, error, dpanic, panic, fatal (default is debug)
	// LOG_LEVEL = os.Getenv("LOG_LEVEL")

	// MYSQL connection dsn
	// ref https://gorm.io/docs/connecting_to_the_database.html#MySQL
	// format user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	MSYQL_DSN = util.LookupEnvStringOr("MSYQL_DSN", "")

	// Redis connection dsn
	// ref https://pkg.go.dev/github.com/go-redis/redis/v8#ParseURL
	// format redis://<user>:<password>@<host>:<port>/<db_number>
	REDIS_DSN = util.LookupEnvStringOr("REDIS_DSN", "")
)
