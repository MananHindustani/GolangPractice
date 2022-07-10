package main

import (
	"github.com/beanstalkd/go-beanstalk"
	"fmt"
	"time"
)
var conn, _ = beanstalk.Dial("tcp", "127.0.0.1:11300")

func main() {
//id, body, err := conn.Reserve(5 * time.Second)

tube := beanstalk.NewTubeSet(conn,"mytube","mytube1")
id, body, err := tube.Reserve(10 * time.Hour)
//id, body, err := conn.Reserve(5 * time.Second)
if err != nil {
panic(err)
}
fmt.Println("job", id)
fmt.Println(string(body))
}
