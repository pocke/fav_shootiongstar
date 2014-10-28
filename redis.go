package main

import (
	"github.com/garyburd/redigo/redis"
)

func userExists(id string) bool {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}

	i, err := redis.Uint64(c.Do("HLEN", id))
	if err != nil {
		panic(err)
	}

	return i != 0
}
