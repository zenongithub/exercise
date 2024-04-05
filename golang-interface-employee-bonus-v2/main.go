package main

import "fmt"

type Employee interface {
	GetBonus() float64
}

type Junior struct {
	Name         string
	BaseSalary   int
	WorkingMonth int
}

type Senior struct {
	Name            string
	BaseSalary      int
	WorkingMonth    int
	PerformanceRate float64
}

type Manager struct {
	Name             string
	BaseSalary       int
	WorkingMonth     int
	PerformanceRate  float64
	BonusManagerRate float64
}

func (j Junior) GetBonus() float64 {
	prorata := float64(j.WorkingMonth) / 12.0
	if prorata > 1 {
		prorata = 1
	}
	return float64(j.BaseSalary) * prorata
}

func (s Senior) GetBonus() float64 {
	prorata := float64(s.WorkingMonth) / 12.0
	if prorata > 1 {
		prorata = 1
	}
	return (2 * float64(s.BaseSalary) * prorata) + (s.PerformanceRate * float64(s.BaseSalary))
}

func (m Manager) GetBonus() float64 {
	prorata := float64(m.WorkingMonth) / 12.0
	if prorata > 1 {
		prorata = 1
	}
	return (2 * float64(m.BaseSalary) * prorata) + (m.PerformanceRate * float64(m.BaseSalary)) + (m.BonusManagerRate * float64(m.BaseSalary))
}

func EmployeeBonus(employee Employee) float64 {
	return employee.GetBonus()
}

func TotalEmployeeBonus(employees []Employee) float64 {
	totalBonus := 0.0
	for _, emp := range employees {
		totalBonus += emp.GetBonus()
	}
	return totalBonus
}

func main() {

	junior := Junior{Name: "Junior 1", BaseSalary: 100000, WorkingMonth: 12}
	fmt.Println("Output for Junior:", EmployeeBonus(junior))

	senior := Senior{Name: "Senior 1", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5}
	fmt.Println("Output for Senior:", EmployeeBonus(senior))

	manager := Manager{Name: "Manager 1", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5, BonusManagerRate: 0.5}
	fmt.Println("Output for Manager:", EmployeeBonus(manager))

	juniorA := Junior{Name: "Junior A", BaseSalary: 100000, WorkingMonth: 12}
	juniorB := Junior{Name: "Junior B", BaseSalary: 100000, WorkingMonth: 12}
	juniorC := Junior{Name: "Junior C", BaseSalary: 100000, WorkingMonth: 12}
	juniors := []Employee{juniorA, juniorB, juniorC}
	fmt.Println("Total Bonus for Juniors:", TotalEmployeeBonus(juniors))

	seniorA := Senior{Name: "Senior A", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5}
	seniorB := Senior{Name: "Senior B", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5}
	seniors := []Employee{seniorA, seniorB}
	fmt.Println("Total Bonus for Seniors:", TotalEmployeeBonus(seniors))

	juniorA = Junior{Name: "Junior A", BaseSalary: 100000, WorkingMonth: 12}
	seniorA = Senior{Name: "Senior A", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5}
	managerA := Manager{Name: "Manager A", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5, BonusManagerRate: 0.5}
	employees := []Employee{juniorA, seniorA, managerA}
	fmt.Println("Total Bonus for All Employees:", TotalEmployeeBonus(employees))
}
