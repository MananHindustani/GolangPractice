package main

import (
	"fmt"
	"time"

	"github.com/beanstalkd/go-beanstalk"
)

func main() {
	c, err := beanstalk.Dial("tcp", "127.0.0.1:11349")
	t1 := beanstalk.NewTube(c, "Tube1")
	t2 := beanstalk.NewTube(c, "Tube2")

	id1, err := t1.Put([]byte("hello,How Are You"), 1, 0, 5*time.Second)

	if err != nil {
		panic(err)
		fmt.Print(err)
	}

	id2, err := t2.Put([]byte("hello,How Are Youuuuuuuuuuuu"), 1, 0, 5*time.Second)

	if err != nil {
		panic(err)
		fmt.Print(err)
	}
	fmt.Println("job:", id1)
	fmt.Println("job:", id2)

}
