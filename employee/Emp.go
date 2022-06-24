package employee

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type employee struct {
	name        string
	dob         string
	yoe         int
	age         int
	salaryBonus int
}

func (e employee) greet() {
	if strings.TrimSpace(e.name) == "" {
		log.Println("Empty name")
	} else {
		if e.yoe <= 0 {
			log.Println("invalid yoe")
		} else {
			fmt.Println("Hello "+e.name+", congratulations for completing", e.yoe, "years")
		}
	}
}
func (e *employee) calcAge() {
	if strings.TrimSpace(e.dob) == "" {
		log.Println("invalid dob")
	} else {
		date := strings.Split(e.dob, "/")
		year, _ := strconv.Atoi(date[2])
		//month, _ := strconv.Atoi(date[1])
		//day, _ := strconv.Atoi(date[0])
		e.age = 2022 - year
	}
}
func (e *employee) calcSalaryBonus() {
	if e.yoe == 1 {
		e.salaryBonus = 20000
	} else if e.yoe >= 2 && e.yoe < 5 {
		e.salaryBonus = 50000
	} else if e.yoe >= 5 && e.yoe <= 10 {
		e.salaryBonus = 100000
	} else if e.yoe > 10 {
		e.salaryBonus = 200000
	} else if e.yoe <= 0 {
		log.Println("invalid yoe")
	}
}
func emp(e employee) employee {
	e.greet()
	e.calcAge()
	e.calcSalaryBonus()
	return e
}
