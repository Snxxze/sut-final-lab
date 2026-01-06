package test

import (
	"sut-final-lab/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestSalary(t *testing.T) {
	g := NewWithT(t)

	t.Run(`Salary is not range 15000 - 200000`, func(t *testing.T) {
		emp := entity.Employees {
			Name: "test",
			Salary: 100,
			EmployeeCode: "HR-1024",
		}

		ok, err := govalidator.ValidateStruct(emp)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
	})
}