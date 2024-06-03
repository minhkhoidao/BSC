package cache

import (
	"backend-speaker-clone/internal/configs"
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

type DataCache struct {
	Key            string
	JsonStringData string
	Tll            time.Duration
}

var redisdb *redis.Client

func MyRedisConnect() {
	log.Println("redis connect")
	redisOption := &redis.Options{
		Addr:     configs.GetRedisHost() + ":" + configs.GetRedisPort(),
		Password: "",
		DB:       configs.GetRedisDb(), // use default DB
	}
	if configs.GetRedisInsecuSkipVerify() {
		redisOption.TLSConfig = &tls.Config{InsecureSkipVerify: configs.GetRedisInsecuSkipVerify()}
	}
	redisdb = redis.NewClient(redisOption)
	data := redisdb.Ping()
	if data.Err() != nil {
		panic(fmt.Sprintf("failed to connect database @ %s", data.Err()))
	}
	// pong, err := redisdb.Ping().Result()
	// log.Println("pong: "+pong, "redis error: "+err.Error())
}

//	func GetRedis() *redis.Client {
//		return redisdb
//	}
//
// ttl : second
func SetCache(key string, value interface{}, ttl int64) error {
	if redisdb == nil {
		return errors.New("redisdb is not connected")
	}

	return redisdb.Set(key, value, time.Duration(ttl*int64(time.Second))).Err()
}

func GetCache(key string) (string, error) {
	if redisdb == nil {
		return "", errors.New("redisdb is not connected")
	}
	return redisdb.Get(key).Result()
}

func IncreaseFlagCounter(key string) (int, error) {
	if redisdb == nil {
		return 0, errors.New("redisdb is not connected")
	}
	result, err := redisdb.Incr(key).Result()
	return int(result), err
}

func DecreaseFlagCounter(key string) (int, error) {
	if redisdb == nil {
		return 0, errors.New("redisdb is not connected")
	}
	result, err := redisdb.Decr(key).Result()
	return int(result), err
}

func Keys(pattern string) ([]string, error) {
	if redisdb == nil {
		return []string{}, errors.New("redisdb is not connected")
	}

	return redisdb.Keys(pattern).Result()
}

// =====================================
// PUSH and POP redis

func RPush(key string, value interface{}) (int64, error) {
	if redisdb == nil {
		return 0, errors.New("redisdb is not connected")
	}
	strCmd := redisdb.RPush(key, value)

	return strCmd.Result()
}

func LPop(key string) (string, error) {
	if redisdb == nil {
		return "", errors.New("redisdb is not connected")
	}
	strCmd := redisdb.LPop(key)
	// log.Println(strCmd.Result())

	return strCmd.Result()
}

func LTrim(key string, start, end int64) (string, error) {
	if redisdb == nil {
		return "", errors.New("redisdb is not connected")
	}
	strCmd := redisdb.LTrim(key, start, end)

	return strCmd.Result()
}

func LRange(key string, start, stop int64) ([]string, error) {
	if redisdb == nil {
		return []string{}, errors.New("redisdb is not connected")
	}
	strCmd := redisdb.LRange(key, start, stop)

	return strCmd.Result()
}

func DelCacheByKey(keys ...string) error {
	if redisdb == nil {
		return errors.New("redisdb is not connected")
	}
	redisdb.Del(keys...)
	return nil
}

func ExpireByKey(key string, ttl int) error {
	if redisdb == nil {
		return errors.New("redisdb is not connected")
	}
	return redisdb.Expire(key, time.Duration(ttl*1000000000)).Err()
}
