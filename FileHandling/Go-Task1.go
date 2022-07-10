package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/sharmamanan796/FileHandling/Config"
)

func readData(chnl chan []string) {
	fmt.Println("read data")

	f, err := excelize.OpenFile("exceldata.xlsx")
	if err != nil {
		fmt.Println("open file error data")

		fmt.Println(err.Error())
		return
	}
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("sheet")
	if err != nil {
		fmt.Println("row error data")

		fmt.Println(err.Error())
		return
	}
	for _, row := range rows {
		fmt.Println("read data")
		a := make([]string, 4)
		for i, colCell := range row {
			fmt.Print("Hello", colCell)
			a[i] = colCell
		}
		chnl <- a
	}
	close(chnl)
}

func writeData(data []string, done chan bool) {
	fmt.Println("write data")

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
	ch := make(chan []string)
	data := make([]string, 4)
	done := make(chan bool)
	go readData(ch)
	for data = range ch {
		//		token := strings.Fields(data)
		go writeData(data, done)
		<-done
		fmt.Println()
	}
}
