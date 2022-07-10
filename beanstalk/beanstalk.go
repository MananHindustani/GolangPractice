package main

import (
	"github.com/beanstalkd/go-beanstalk"
	"fmt"
	"time"
)
var conn, _ = beanstalk.Dial("tcp", "127.0.0.1:11300")

func main() {
	fmt.Println("Beanstalk")
//	id, err := conn.Put([]byte("myjob"), 1, 0, time.Minute)
	tube := beanstalk.NewTube(conn, "mytube")
	tube1 := beanstalk.NewTube(conn, "mytube1")

	id, err := tube.Put([]byte("myjob"), 1, 0, time.Minute)
	id1, err := tube1.Put([]byte("myjob1"), 1, 0, time.Minute)

	if err != nil 	{
			panic(err)
	}
	fmt.Println("job", id)
	fmt.Println("job", id1)

}