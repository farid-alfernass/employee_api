package cache

import (
	"sync"

	"github.com/faridtriwicaksono/employee_api/infrastructure/cache/redis"
)

var (
	redisConnOnce sync.Once
	redisConn     redis.Cache
)

// GetRedisAccess from go-redis
func GetRedisAccess() redis.Cache {
	if redisConn != nil {
		return redisConn
	}

	redisConnOnce.Do(func() {
		redisClient := redis.GetRedis()
		redisConn = redisClient
	})

	return redisConn
}
