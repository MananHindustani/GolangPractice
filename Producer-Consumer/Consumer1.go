package main

import (
	"fmt"
	"time"

	"github.com/beanstalkd/go-beanstalk"
)

func main() {
	// stop := make(chan bool)

	c, _ := beanstalk.Dial("tcp", "127.0.0.1:11349")
	ts := beanstalk.NewTubeSet(c, "Tube1", "Tube2")
	// ts1 := beanstalk.NewTubeSet(c, "Tube2")
	i := 1
	for i <= 4 {
		id, body, err := ts.Reserve(10 * time.Hour)
		if err != nil {
			panic(err)
		}
		fmt.Println("job", id)
		fmt.Println(string(body))
		i++
	}
	// id, body, err = ts1.Reserve(10 * time.Hour)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("job", id)
	// fmt.Println(string(body))

	// <-stop

}
