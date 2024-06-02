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

// ref https://github.com/redis/go-redis?tab=readme-ov-file#quickstart

package redis

import (
	"sync"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/xiexianbin/go-echo-demo/config"
	"github.com/xiexianbin/go-echo-demo/pkg/util"
)

var (
	rdb  *redis.Client
	once sync.Once
)

func Init() {
	once.Do(func() {
		redis_dsn := config.REDIS_DSN
		if redis_dsn != "" {
			opt, err := redis.ParseURL(redis_dsn)
			util.Mustf(err, "init Redis failed")

			opt.DialTimeout = 3 * time.Second // no time unit = seconds
			opt.ReadTimeout = 6 * time.Second

			// connection pool
			opt.PoolFIFO = true
			opt.PoolSize = 10

			rdb = redis.NewClient(opt)
		}
	})
}

func Get(key string) ([]byte, error) {

	return nil, nil
}
