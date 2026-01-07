package entity

import "gorm.io/gorm"

type Employees struct {
	gorm.Model
	Name		string		`valid:"stringlength(2|80)~Name must len 2-80"`
	Salary		float64		`valid:"range(150000|200000)~Salary must be between 15000 and 200000"`
	EmployeeCode	string	`valid:"matches(^A-Z\\-d{4}$)~EmployeeCode must be 2 uppercase English letters (A-Z) followed by ‘-’ and 4 digits (0-9)"`
}