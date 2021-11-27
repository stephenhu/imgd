package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

const (
	HASH_PRODUCTS				= "imgd.urls.products"
	HASH_IMAGES					= "imgd.urls.images"
	QUEUE_PRODUCTS			= "imgd.queue.products"	
)

const (
	REDIS_DEFAULT       = "127.0.0.1:6379"
)


func redisStr() string {
  return fmt.Sprintf("%s:%s", conf.Redis.Host, conf.Redis.Port)
} // redisStr


func (r Red) Connect() {

	c := redis.NewClient(&redis.Options{
		Addr:					redisStr(),
		DialTimeout: 	10 * time.Second,
		ReadTimeout: 	30 * time.Second,
		WriteTimeout:	30 * time.Second,
		PoolSize:			10,
		PoolTimeout:	30 * time.Second,
	})

	err := c.Ping().Err()

	if err != nil {
		log.Fatalf("Unable to connect to Redis: %s", err.Error())
	}

	client = c

	log.Printf("Redis connection established at %s", redisStr())

} // Connect
