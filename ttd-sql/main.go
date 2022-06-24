package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type employee struct {
	empId     int
	empName   string
	empPos    string
	empExp    float64
	empSalary float64
}

var dbName, dbSource = "mysql", "root:Goyal@921@tcp(127.0.0.1:3306)/test"

func post(emp employee) (employee, error) {
	db, err := sql.Open(dbName, dbSource)
	if err != nil {
		log.Print(err)
	}
	defer db.Close()
	if emp.empId <= 0 {
		return employee{}, errors.New("Invalid Id")
	}
	query := fmt.Sprintf("INSERT INTO Employee (emp_id, emp_name, emp_position, emp_exp, emp_salary)\nVALUES\n(%v,'%v','%v',%v,'%v');", emp.empId, emp.empName, emp.empPos, emp.empExp, emp.empSalary)
	data, err := db.Exec(query)
	if err != nil {
		return employee{}, errors.New("Emp Already Exists")
	}
	fmt.Println(data)
	return emp, nil
}

func get(empId int) (employee, error) {
	db, err := sql.Open(dbName, dbSource)
	if err != nil {
		log.Print(err)
	}
	defer db.Close()
	Query := fmt.Sprintf("select * from Employee where emp_id=%v", empId)
	data := db.QueryRow(Query)
	outEmp := employee{}
	if err := data.Scan(&outEmp.empId, &outEmp.empName, &outEmp.empPos, &outEmp.empExp, &outEmp.empSalary); err != nil {
		if err == sql.ErrNoRows {
			return employee{}, errors.New("Emp does not exist")
		}
	}
	return outEmp, nil
}
