package student

import "fmt"

var student = make(map[string]string)

func SetstudentDetails() {

	student["First Name"] = "Manan"
	student["Last Name"] = "Manan"
	student["Course"] = "MCA"
	student["Year"] = "Third"
}

func GetstudentDetails() {
	for key, value := range student {
		fmt.Printf("student[%s] = %s\n", key, value)
	}

}
