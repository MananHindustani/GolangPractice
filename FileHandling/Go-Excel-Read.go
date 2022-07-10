package main

import "github.com/360EntSecGroup-Skylar/excelize"

func main() {
	f, err := excelize.OpenFile("Hello.xlsx")
	if err != nil {
		println(err.Error())
		return

	}
	// Get value from cell by given worksheet name and axis.
	/*	cell, err := f.GetCellValue("Sheet1", "B2")

		if err != nil {

			println(err.Error())

			return

		}

		println(cell)
	*/
	// Get all the rows in the Sheet1.

	rows, err := f.GetRows("Sheet1")

	for _, row := range rows {

		for _, colCell := range row {

			print(colCell, "\t")

		}

		println()

	}

}
