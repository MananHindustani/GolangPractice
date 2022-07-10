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
	/*
		//SET Key

		rdb.Set(ctx, "gokey", "golang", 0).Err()
		if err != nil {
			panic(err)
		}
		//GET Key
		rdb.Get(ctx, "mygokey1")
		fmt.Println()
	*/

	/*	//SET only if it eXists

			n := rdb.SetXX(ctx, "mygokey1", "goal", 0)

			fmt.Println(n)

			//SET only if Not does not already eXists

			n = rdb.SetNX(ctx, "mygokey1", "goal", 0)

			fmt.Println(n)

		fmt.Println(rdb.Get(ctx, "mygokey1"))
	*/
	//MSETX
	/*rdb.MSet(ctx, "Uone", "1", "Utwo", "2", "Uthree", "3")
	n := rdb.MSetNX(ctx, "Ufour", "4")
	fmt.Println(n)

	fmt.Println(rdb.MGet(ctx, "Uone", "Utwo", "Uthree", "Ufour"))
	*/
	/*
		//GETRANGE
		s := rdb.GetRange(ctx, "gokey", 0, 3)
		fmt.Println(s)
	*/

	/*
		//INCREMENT DECREMENT

		fmt.Println(rdb.Incr(ctx, "Uone"))
		fmt.Println(rdb.Decr(ctx, "Uone"))

		fmt.Println(rdb.IncrBy(ctx, "Ufour", 4))
		fmt.Println(rdb.DecrBy(ctx, "Ufour", 4))

		fmt.Println(rdb.IncrByFloat(ctx, "Uthree", 3.1))
	*/
	/*	//Append
		rdb.Append(ctx, "gokey", " iDE")
		fmt.Println(rdb.Get(ctx, "gokey"))
	*/
	//BIT Operation
	fmt.Println(rdb.Get(ctx, "mygokey"))
	fmt.Println(rdb.Get(ctx, "mygokey1"))

	init1 := rdb.BitOpAnd(ctx, "result1", "mygokey", "mygokey1")
	fmt.Println(init1)
	fmt.Println(rdb.Get(ctx, "result1"))
	//	DOUBT
	init2 := rdb.BitOpOr(ctx, "result2", "mygokey", "mygokey1")
	fmt.Println(init2)
	fmt.Println(rdb.Get(ctx, "result2"))

	//	rdb.BitCount{0, 4}
	//	rdb.BitCount(ctx, "gokey")

}
