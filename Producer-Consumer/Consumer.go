package main

import (
	"fmt"
	"time"

	"github.com/beanstalkd/go-beanstalk"
)

func main() {
	c, err := beanstalk.Dial("tcp", "127.0.0.1:11349")
	id, body, err := c.Reserve(10 * time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Println(id, body)
}
