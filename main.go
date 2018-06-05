package main

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"fmt"
	"github.com/redis-wrapper"
)

type Client struct {
	RedisClient *redis.Client
}

func main() {
	c := redis_wrapper.Client{}
	client, err := c.Create()
	if err != nil {
		logrus.Error(err.Error())
	}
	var keys []string
	var cursor uint64
	keys, cursor, err = client.RedisClient.Scan(cursor, "", 100).Result()
	for _, k := range keys {
		val, err := client.Get(k)
		if err != nil {
			logrus.Error(err.Error())
		}
		fmt.Printf("%v \n", val)
	}
}
