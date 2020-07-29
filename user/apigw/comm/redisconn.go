package comm

import (
	"github.com/go-redis/redis/v8"
)

//var ctx = context.Background()

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	//尝试链接情况
	// pong, err := rdb.Ping(ctx).Result()
	// fmt.Println(pong, err)
	// Output: PONG <nil>
	return rdb
}