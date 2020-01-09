package beego

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
	"github.com/vmihailenco/msgpack"
	"strconv"
	"time"

	//Redis
	_ "github.com/astaxie/beego/cache/redis"
)

var (
	// Red ...
	Red, redisErr = cache.NewCache("redis", beego.AppConfig.String("redisconn"))
)

// IsRedisOk ...
func IsRedisOk() bool {
	if redisErr != nil {
		fmt.Println(redisErr.Error())
	}
	return redisErr == nil
}

func RememberData(key string, timeout time.Duration, result interface{}, callFun func() (interface{}, error)) error {
	if IsRedisOk() && Red.IsExist(key) {
		data := Red.Get(key).([]uint8)
		unPackErr := msgpack.Unmarshal(data, result)
		if unPackErr == nil {
			logs.Info("load from redis " + key)
			return nil
		}
	}
	funResult, err := callFun()
	if err != nil {
		return err
	}
	packData, strErr := msgpack.Marshal(funResult)
	if strErr == nil && IsRedisOk() {
		go Red.Put(key, string(packData), timeout)
	}
	return err
}

// GetRedisInt ...
func GetRedisInt(key string, def int) int {
	if !IsRedisOk() {
		return def
	}
	if Red.IsExist(key) {
		numStr := string(Red.Get(key).([]uint8))
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return def
		}
		return num
	}
	return def
}

// GetRedisString ...
func GetRedisString(key string, def string) string {
	if !IsRedisOk() {
		return def
	}
	if Red.IsExist(key) {
		return string(Red.Get(key).([]uint8))
	}
	return def
}

func GetIsExist(key string) bool {
	if !IsRedisOk() {
		return false
	}
	return Red.IsExist(key)
}

// IsRedisExist ...
func IsRedisExist(key string) bool {
	if !IsRedisOk() {
		return false
	}
	return Red.IsExist(key)
}

// IncrRedisKey ...
func IncrRedisKey(key string) {
	if IsRedisOk() {
		if Red.IsExist(key) {
			Red.Incr(key)
		} else {
			Red.Put(key, 1, 14*24*time.Hour)
		}
	}
}

// PutRedis
func PutRedis(key string, val interface{}, timeout time.Duration) {
	if IsRedisOk() {
		Red.Put(key, val, timeout)
	}
}

func DeleteRedis(key string) {
	if IsRedisOk() {
		Red.Delete(key)
	}
}
