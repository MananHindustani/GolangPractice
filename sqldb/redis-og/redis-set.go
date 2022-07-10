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
	//Create a set
	/*
		xrdb.SAdd(ctx, "goset", "one", "two", "three", "four")
		setsresult := rdb.SMembers(ctx, "goset")
		fmt.Println(setsresult)


		rdb.SAdd(ctx, "goset1", "three", "six", "four", "five")

	*/
	setsresult1 := rdb.SMembers(ctx, "goset")
	fmt.Println(setsresult1)

	setsresult2 := rdb.SMembers(ctx, "goset1")
	fmt.Println(setsresult2)

	n1 := rdb.SCard(ctx, "goset")
	n2 := rdb.SCard(ctx, "goset1")

	fmt.Println(n1)
	fmt.Println(n2)
	/*
		//	set difference , Intersection and Union

		//set difference
		setdiff := rdb.SDiff(ctx, "goset", "goset1")
		fmt.Println(setdiff)
		//set difference stored
		rdb.SDiffStore(ctx, "resultset", "goset", "goset1")
		fmt.Println(rdb.SMembers(ctx, "resultset"))

		//set intersection
		setinter := rdb.SInter(ctx, "goset", "goset1")
		fmt.Println(setinter)

		//set intersection stored
		rdb.SInterStore(ctx, "resultinterset", "goset", "goset1")

		fmt.Println(rdb.SMembers(ctx, "resultinterset"))
	*/

	/*
		//to check wether given element is member or not
		isSetMember := rdb.SIsMember(ctx, "goset", "nine")
		fmt.Println(isSetMember)
	*/
	/*
		//pop and remove fn

		spop := rdb.SPop(ctx, "goset")
		fmt.Println(spop)

		snpop := rdb.SPopN(ctx, "goset1", 2)
		fmt.Println(snpop)

		//remove
		srem := rdb.SRem(ctx, "goset1", "four")
		fmt.Println(srem)
	*/
	sunion := rdb.SUnion(ctx, "goset", "goset1")
	fmt.Println(sunion)

	rdb.SUnionStore(ctx, "sunionstore", "goset", "goset1")
	fmt.Println(rdb.SMembers(ctx, "sunionstore"))

	/*
		//Random Numbers

		setrandom := rdb.SRandMember(ctx, "goset1")
		fmt.Println(setrandom)

		setrandomN := rdb.SRandMemberN(ctx, "goset1", 2)
		fmt.Println(setrandomN)
	*/
	//	fmt.Println(rdb.SScan(ctx, "goset", 2, "one", 3))

}
