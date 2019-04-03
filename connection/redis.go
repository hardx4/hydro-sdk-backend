package connection

import (
	"github.com/HydroProtocol/hydro-sdk-backend/utils"
	"github.com/go-redis/redis"
)

func NewRedisClient(url string) *redis.Client {
	opt, err := redis.ParseURL(url)

	if err != nil {
		panic(err)
	}

	// Redis supports many configs.
	// You should change them by demand.
	opt.PoolSize = 10
	opt.MaxRetries = 2

	client := redis.NewClient(opt)
	utils.Info("Redis Connect")
	return client
}
