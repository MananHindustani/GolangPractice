package main

import (
	"fmt"
	"time"

	"github.com/beanstalkd/go-beanstalk"
)

func producer(c *beanstalk.Conn, ch chan bool) {
	t1 := beanstalk.NewTube(c, "Tube1")
	t2 := beanstalk.NewTube(c, "Tube2")

	id1, err := t1.Put([]byte("hello,How Are Youuuuuuuuuuuu"), 1, 0, 120*time.Second)

	if err != nil {
		panic(err)
		fmt.Print(err)
	}
	fmt.Println("job:", id1)
	/*	id2, err := t2.Put([]byte("hello,How Are Youuuuuuuuuuuu"), 1, 0, 120*time.Second)

		if err != nil {
			panic(err)
			fmt.Print(err)
		}

		fmt.Println("job:", id2)*/
	ch <- true
}

func consumer(c *beanstalk.Conn, ch chan bool) {
	ts := beanstalk.NewTubeSet(c, "Tube1", "Tube2")
	id, body, err := ts.Reserve(10 * time.Hour)
	if err != nil {
		panic(err)
	}
	fmt.Println("job", id)
	fmt.Println(string(body))
	/*id, body, err := c.Reserve(10 * time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Println("job", id)
	fmt.Println(string(body)) */
	ch <- true
}
func main() {
	ch := make(chan bool)
	c, err := beanstalk.Dial("tcp", "127.0.0.1:11349")
	if err != nil {
		panic(err)
		fmt.Print(err)
		ch <- true

	}

	go producer(c, ch)
	<-ch
	go consumer(c, ch)
	<-ch
}
