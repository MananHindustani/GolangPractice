package main

import (
	"fmt"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/beanstalkd/go-beanstalk"
)

func read(line chan string) {

	f, err := excelize.OpenFile("Hello.xlsx")
	if err != nil {
		println(err.Error())
		return

	}
	// Get all the rows in the Sheet1.

	rows, err := f.GetRows("Sheet1")
	for _, row := range rows {
		cell := ""
		for _, colCell := range row {
			print(colCell, "\t")
			cell = cell + colCell + " "
		}
		line <- cell
		println()
	}
	close(line)
}

func producer(c *beanstalk.Conn, ch chan bool) {
	t1 := beanstalk.NewTube(c, "Tube1")
	s := make([]string, 5)
	i := 0
	line := make(chan string)
	go read(line)
	for row := range line {
		fmt.Println(i)

		s[i] = row
		id, err := t1.Put([]byte(s[i]), 1, 0, 120*time.Second)
		i++
		if err != nil {
			panic(err)
			fmt.Print(err)

		}
		fmt.Println("job:", id)
	}

	ch <- true
}

func consumer(c *beanstalk.Conn, ch chan bool) {
	ts := beanstalk.NewTubeSet(c, "Tube1")
	id, body, err := ts.Reserve(10 * time.Hour)
	if err != nil {
		panic(err)
	}
	fmt.Println("job", id)
	fmt.Println(string(body))
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
