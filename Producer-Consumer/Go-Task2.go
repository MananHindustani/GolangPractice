package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sharmamanan796/FileHandling/Config"
)

func readData(chnl chan string) {
	db, err := Config.GetConnection()
	defer db.Close()
	rows, err := db.Query("select * from student")
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {

		var id int
		var name string
		var email string
		var yearofjoin string
		err := rows.Scan(&id, &name, &email, &yearofjoin)
		if err != nil {
			panic(err)
		}
		row := strconv.Itoa(id) + "\t\t" + name + "\t\t" + email + "\t\t" + yearofjoin
		//		fmt.Printf("%s", row)
		chnl <- row
	}
	close(chnl)
}

func writeData(f *os.File, row string, isDone chan bool) {

	fmt.Fprintln(f, row)
	isDone <- true

}
func main() {
	ch := make(chan string)
	isDone := make(chan bool)
	f, err := os.Create("text3.txt")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	go readData(ch)
	for row := range ch {
		fmt.Println(row)
		go writeData(f, row, isDone)
		<-isDone
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	fmt.Println("file written successfully")

}
