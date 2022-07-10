package main

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	/*
						//set feilds and values

		rdb.HSet(ctx, "gohashstates", "MP", "Bhopal", "CG", "Raipur", "RJ", "Jaipur")
						//Set multiple values
		rdb.HMSet(ctx, "gohashstates", "Bihar", "PAtna", "J&K", "Shrinagar")

						//set if not exist
		fmt.Println(setNx)

	*/

	/*
			//get feild of key
		hGet := rdb.HGet(ctx, "gohashstates", "RJ")
		fmt.Println(hGet)
			//get Multiple vlues respective feilds
		hMGet := rdb.HMGet(ctx, "gohashstates", "RJ", "MP", "MH")
		fmt.Println(hMGet)

			//get all feild-values  of specified key
		hGetAll := rdb.HGetAll(ctx, "gohashstates")
		fmt.Println(hGetAll)

	*/

	/*
				//All feilds and values of specified key
		hashkeys := rdb.HKeys(ctx, "gohashstates")
		fmt.Println(hashkeys)

		hashvals := rdb.HVals(ctx, "gohashstates")
		fmt.Println(hashvals)
	*/
	/*
		//check existence of existence feild
		hexist := rdb.HExists(ctx, "gohashstates", "JH")
		fmt.Println(hexist)
	*/
	/*
		//delete feild
		hdel := rdb.HDel(ctx, "gohashstates", "JH")
		fmt.Println(hdel)
	*/

	/*
		//length of hash key
		hlen := rdb.HLen(ctx, "gohashstates")
		fmt.Println(hlen)
	*/

	/*
			//increment by int
		incrBy := rdb.HIncrBy(ctx, "myhashes", "f6", 10)
		fmt.Println(incrBy)
			//increment by floatm
		incrByFloat := rdb.HIncrByFloat(ctx, "myhashes", "f6", 10.10)
		fmt.Println(incrByFloat)

	*/
	/*
		hscan := rdb.HScan(ctx, "gohashstates", 0, "MP", 2)
		fmt.Println(hscan)
	*/
}
