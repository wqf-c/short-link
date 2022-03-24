package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
	"short-link/common"
	"time"
)

var ctx = context.Background()
var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     common.Info.Redis.Ip + ":" + common.Info.Redis.Port,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

}

func RedisSet(key string, value interface{}, expiration time.Duration) error {

	err := rdb.Set(ctx, key, value, expiration).Err()

	return err
}

func RedisGet(key string) (string, error) {

	val, err := rdb.Get(ctx, key).Result()

	return val, err
}
