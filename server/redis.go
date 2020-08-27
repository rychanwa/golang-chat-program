package main

import (

	"github.com/gomodule/redigo/redis"
	"time"
)

var pool *redis.Pool

func initPool(addr string) {
	pool = &redis.Pool{
		MaxIdle: 10,
		MaxActive: 0,
		IdleTimeout: 10 * time.Second,
		Dial: func () (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
	}
}
