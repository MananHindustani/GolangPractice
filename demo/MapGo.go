package main

import "fmt"

func main() {

	student := make(map[string]string)

	student["First Name"] = "Manan"
	student["Last Name"] = "Manan"
	student["Course"] = "MCA"
	student["Year"] = "Third"

	for key, value := range student {
		fmt.Printf("student[%s] = %s\n", key, value)
	}

	/*employeeSal := make(map[string]int)
	employeeSal["Manan"] = 5000
	employeeSal["Rupal"] = 5000
	employeeSal["Hari"] = 5000
	emplyee := "Manan"
	salary := employeeSal[emplyee]
	fmt.Println(employeeSal)
	fmt.Println(salary)
	//	value, ok := employeeSal["jules"]

	for key, value := range employeeSal YEA{
		fmt.Printf("employeeSalary[%s] = %d\n", key, value)
	}
	delete(employeeSal, "Hari")
	fmt.Println(employeeSal)*/
}
