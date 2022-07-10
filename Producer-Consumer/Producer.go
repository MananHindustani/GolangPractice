package main

import (
	"fmt"
	"time"

	"github.com/beanstalkd/go-beanstalk"
)

func main() {
	c, err := beanstalk.Dial("tcp", "127.0.0.1:11349")

	id, err := c.Put([]byte("hello,Bro"), 1, 0, 120*time.Second)
	if err != nil {
		panic(err)
		fmt.Print(err)
	}
	fmt.Println(id)
}
