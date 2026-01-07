package test

import (
	"sut-final-lab/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestEmpCode(t *testing.T) { 
	g := NewWithT(t) 
	
	t.Run(`EmployeeCode is not format`, func(t *testing.T) { 
		employee := entity.Employees { 
			Name: "John Doe", 
			Salary: 20000.00, 
			EmployeeCode: "HR102400", 
		}

		ok, err := govalidator.ValidateStruct(employee) 
		
		g.Expect(ok).NotTo(BeTrue()) 
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("EmployeeCode must be 2 uppercase English letters (A-Z) followed by - and 4 digits (0-9)"))
	}) 	
}