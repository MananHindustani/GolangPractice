package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong)
	//	rdb.ZAdd(ctx, "gosorted", rdb.Z{1, "one"})
	sortedcard := rdb.ZCard(ctx, "mysorted")
	fmt.Println(sortedcard)
	sortedRange := rdb.ZRange(ctx, "mysorted", 0, -1)
	fmt.Println(sortedRange)
}
