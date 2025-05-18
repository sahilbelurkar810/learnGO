package main

import (
	"fmt"
)

type Employee interface {
	GetName() string
	GetSalary() float64
	GetRole() string
}

type Contractor interface {
	GetName() string
	GetRate() float64
	GetHours() int
}

// Implement Employee interface
type FullTimeEmployee struct {
	name   string
	salary float64
	role   string
}

func (e FullTimeEmployee) GetName() string {
	return e.name
}

func (e FullTimeEmployee) GetSalary() float64 {
	return e.salary
}

func (e FullTimeEmployee) GetRole() string {
	return e.role
}

// Implement Contractor interface
type ContractWorker struct {
	name  string
	rate  float64
	hours int
}

func (c ContractWorker) GetName() string {
	return c.name
}

func (c ContractWorker) GetRate() float64 {
	return c.rate
}

func (c ContractWorker) GetHours() int {
	return c.hours
}

func main() {
	// Create a full time employee
	employee := FullTimeEmployee{
		name:   "John Doe",
		salary: 75000.00,
		role:   "Software Engineer",
	}

	// Create a contractor
	contractor := ContractWorker{
		name:  "Jane Smith",
		rate:  100.00,
		hours: 40,
	}

	// Print employee details
	fmt.Printf("Employee: %s, Role: %s, Salary: $%.2f\n",
		employee.GetName(), employee.GetRole(), employee.GetSalary())

	// Print contractor details
	fmt.Printf("Contractor: %s, Rate: $%.2f, Hours: %d\n",
		contractor.GetName(), contractor.GetRate(), contractor.GetHours())
}
