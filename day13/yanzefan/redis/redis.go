package redis

import (
	"flag"
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	redisAddress string
	pool         *redis.Pool
)

func init() {
	flag.StringVar(&redisAddress, "s", "127.0.0.1:6379", "redis server address")
	flag.Parse()

	poolSize := 20
	pool = &redis.Pool{
		MaxIdle:     poolSize,
		IdleTimeout: time.Minute,
		Dial: func() (conn redis.Conn, err error) {
			conn, err = redis.Dial("tcp", redisAddress)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
}

func GetRedisConn() redis.Conn {
	return pool.Get()
}

