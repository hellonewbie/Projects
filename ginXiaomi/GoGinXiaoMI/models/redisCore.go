package models

import (
	"context"
	"fmt"
	"ginXiaomi/GoGinXiaoMI/utils"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
)

var redisC *RedisClient
var redisCoretxt = context.Background()
var RedisC1 *redis.Client

type RedisClient struct {
	Addr     string
	Password string
	Db       int
}

func (rc *RedisClient) SetRedisAddr(addr string) func(*RedisClient) interface{} {
	return func(rc *RedisClient) interface{} {
		rcA := rc.Addr
		rc.Addr = addr
		return rcA
	}
}

func (rc *RedisClient) SetRedisPwd(pwd string) func(*RedisClient) interface{} {
	return func(rc *RedisClient) interface{} {
		rcp := rc.Password
		rc.Password = pwd
		return rcp
	}
}

func (rc *RedisClient) SetRedisDb(db int) func(*RedisClient) interface{} {
	return func(rc *RedisClient) interface{} {
		rcdb := rc.Db
		rc.Db = db
		return rcdb
	}
}

func (rc *RedisClient) RedisInit(options ...func(*RedisClient) interface{}) (*redis.Client, error) {
	Db, err := strconv.Atoi(utils.RedisDb)
	if err != nil {
		log.Fatal("RedisDb transform into int failed")
	}
	q := &RedisClient{
		Addr:     utils.Addr,
		Password: utils.Password,
		Db:       Db,
	}
	for _, option := range options {
		option(q)
	}
	redisC = q

	client := redis.NewClient(&redis.Options{
		Addr:     redisC.Addr,
		Password: redisC.Password, // no password set
		DB:       redisC.Db,       // use default DB
	})

	_, err = client.Ping(redisCoretxt).Result()

	if err != nil {
		fmt.Println("content", "redis client ping is error", "error", err.Error())
		return nil, err
	}
	RedisC1 = client
	return client, nil
}
