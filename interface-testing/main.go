package main

import "fmt"

type employee struct {
	age  int
	name string
	add  string
}

func validate(emp employee) (bool, employee) {
	if emp.age < 22 {
		emp.age = 0
		emp.name = ""
		emp.add = ""
		return false, emp
	} else {
		return true, emp
	}
}
func main() {
	emp := employee{20, "Tushar", "HSR"}
	aged, emp := validate(emp)
	fmt.Println(aged, emp)
}
