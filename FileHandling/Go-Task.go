package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sharmamanan796/FileHandling/Config"
)

func readData(chnl chan string) {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		//		fmt.Println(s.Text())
		chnl <- s.Text()
	}
	close(chnl)
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func writeData(data []string, done chan bool) {
	db, _ := Config.GetConnection()
	defer db.Close()
	sqlStatement := `
	INSERT INTO student (id, name, email,yearofjoin)
	VALUES ($1, $2, $3,$4)
	RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, data[0], data[1], data[2], data[3]).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)

	done <- true
}

func main() {
	ch := make(chan string)
	go readData(ch)
	done := make(chan bool)
	for data := range ch {
		token := strings.Fields(data)
		go writeData(token, done)
		<-done

		fmt.Println()
}

	}

}
