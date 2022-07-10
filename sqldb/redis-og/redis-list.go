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
	/*
		//push from Left into a List
		listLPush := rdb.LPush(ctx, "golist1", "l1", "l2", "l3", "l4")
		fmt.Println(listLPush)
		//length of list
		n1 := rdb.LLen(ctx, "golist1")
		fmt.Println(n1)

		listRPush := rdb.RPush(ctx, "golist2", "l1", "l2", "l3", "l4")
		fmt.Println(listRPush)
		//length of list
		n2:= rdb.LLen(ctx, "golist1")
		fmt.Println(n2)

	*/

	/*
		isLpush := rdb.LPushX(ctx, "golist2", "l1")

		fmt.Println(isLpush)

		isRpush := rdb.RPushX(ctx, "golist1", "l1")

		fmt.Println(isRpush)

	*/
	/*
		LInsert with three ways

		//Insert function
		lInsert := rdb.LInsert(ctx, "golist2", "BEFORE", "l3", "lala")
		fmt.Println(lInsert)


		rdb.LInsertBefore(ctx, "golist2", "lala", "lalu")


		//LInsertAfter
		rdb.LInsertAfter(ctx, "golist1", "lala", "lalu")


	*/

	/*
			LPOP ,RPOP

		list1 := rdb.LRange(ctx, "golist1", 0, -1)
		fmt.Println(list1)

		rdb.LPop(ctx, "golist1")

		list1 = rdb.LRange(ctx, "golist1", 0, -1)
		fmt.Println(list1)

		list2 := rdb.LRange(ctx, "golist2", 0, -1)
		fmt.Println(list2)

		rdb.RPop(ctx, "golist2")

		list2 = rdb.LRange(ctx, "golist2", 0, -1)
		fmt.Println(list2)
	*/

	/*
		// Remove specified element n times
		rdb.LRem(ctx, "golist2", 2, "lalu")

		//trim then
		list := rdb.LRange(ctx, "golist1", 0, -1)
		fmt.Println(list)
		rdb.LTrim(ctx, "golist1", 2, 3)
		list = rdb.LRange(ctx, "golist1", 0, -1)
		fmt.Println(list)
	*/
	/*
			//positoion and index


		atIndex := rdb.LIndex(ctx, "golist2", 2)
		fmt.Println(atIndex)
	*/
	list := rdb.LSet(ctx, "golist1", 0, "l2")
	fmt.Println(list)

	list1 := rdb.LRange(ctx, "golist1", 0, -1)
	fmt.Println(list1)

}
