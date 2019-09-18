package redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/faridtriwicaksono/employee_api/lib/converter"
)

var (
	mutex sync.Mutex
)

// Cache redis interface
type (
	Cache interface {
		Get(key string) *redis.StringCmd
		Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
		Del(...string) *redis.IntCmd
	}
)

// GetRedis connection
func GetRedis() Cache {
	mutex.Lock()
	defer mutex.Unlock()
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       9,
	})

	return client
}

// GetRedisValue return byte
func GetRedisValue(client Cache, key string) ([]byte, error) {
	if client == nil {
		return nil, errors.New("nil redis pointer")
	}
	res, err := client.Get(key).Result()
	if err != nil || string(res) == "" {
		return nil, err
	}

	data := struct {
		Values json.RawMessage `json:"values"`
		Expire int             `json:"expire"`
	}{}

	err = json.Unmarshal([]byte(res), &data)
	if err != nil {
		return nil, err
	}
	return data.Values, nil
}

// SetRedisValue with Set
func SetRedisValue(client Cache, key string, value interface{}, expireTime int) {
	if client == nil {
		return
	}

	ttl := time.Duration(expireTime) * time.Second
	expireUnix := int(time.Now().Unix())

	valueString := converter.ConvertString(value)

	if valueString == "" {
		valueString = `""`
	}

	redisString := fmt.Sprintf(`{"values" : %s, "expire" : %d}`, valueString, expireUnix)

	client.Set(key, redisString, ttl)
}

// DelRedisValue with Set
func DelRedisValue(client Cache, key string) error {
	if client == nil {
		return errors.New("nil redis pointer")
	}
	res, err := client.Del(key).Result()
	if err != nil || string(res) == "" {
		return err
	}
	fmt.Printf("res del %+v \n", res)

	return nil
}
