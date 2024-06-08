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

// Package redis for redis client
// - github.com/go-redis/redis/v8 for redis 6, https://github.com/redis/go-redis/blob/v8.11.5/Makefile#L19
// - github.com/redis/go-redis/v9 for redis 7, https://github.com/redis/go-redis/blob/v9.5.2/Makefile#L34
// ref https://github.com/redis/go-redis?tab=readme-ov-file#quickstart
package redis

import (
	"context"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/xiexianbin/go-echo-demo/config"
	"github.com/xiexianbin/go-echo-demo/pkg/util"
)

var (
	rdb  *redis.Client
	rdbc *redis.ClusterClient
	once sync.Once
)

func Init() {
	once.Do(func() {
		// Redis Single
		if config.REDIS_DSN != "" {
			opt, err := redis.ParseURL(config.REDIS_DSN)
			util.Mustf(err, "parse Redis by REDIS_DSN failed")

			opt.DialTimeout = 3 * time.Second // no time unit = seconds
			opt.ReadTimeout = 6 * time.Second

			// connection pool
			opt.PoolFIFO = true
			opt.PoolSize = 10

			// https://pkg.go.dev/github.com/redis/go-redis/v9#NewClient
			rdb = redis.NewClient(opt)
			_, err = rdb.Ping(context.TODO()).Result()
			util.Mustf(err, "ping redis err")
			return
		}

		// Redis Sentinel Cluster
		if config.REDIS_SENTINEL_DSN != "" {
			opt, err := redis.ParseClusterURL(config.REDIS_DSN)
			util.Mustf(err, "init Redis by REDIS_SENTINEL_DSN failed")

			// https://pkg.go.dev/github.com/redis/go-redis/v9#NewFailoverClient
			rdb = redis.NewFailoverClient(&redis.FailoverOptions{
				MasterName:    "master",
				ClientName:    opt.ClientName,
				SentinelAddrs: opt.Addrs,
				Password:      opt.Password,

				DialTimeout: 3 * time.Second, // no time unit = seconds
				ReadTimeout: 6 * time.Second,

				PoolFIFO: true,
				PoolSize: 10,
			})

			_, err = rdb.Ping(context.TODO()).Result()
			util.Mustf(err, "ping redis sentinel err")
			return
		}

		// Redis Cluster
		if config.REDIS_CLUTER_DSN != "" {
			opt, err := redis.ParseClusterURL(config.REDIS_CLUTER_DSN)
			util.Mustf(err, "parse Redis by REDIS_CLUTER_DSN failed")

			opt.DialTimeout = 3 * time.Second // no time unit = seconds
			opt.ReadTimeout = 6 * time.Second

			// connection pool
			opt.PoolFIFO = true
			opt.PoolSize = 10

			// https://pkg.go.dev/github.com/redis/go-redis/v9#NewClusterClient
			rdbc = redis.NewClusterClient(opt)
			_, err = rdb.Ping(context.TODO()).Result()
			util.Mustf(err, "ping redis cluster err")
			return
		}
	})
}

func Get(key string) ([]byte, error) {

	return nil, nil
}
