package test

import (
	"sut-final-lab/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestValid(t *testing.T) {
	g := NewWithT(t)

	t.Run(`positive`, func(t *testing.T) {
		employee := entity.Employees{
			Name:         "test",
			Salary:       20000,
			EmployeeCode: "HR-1024",
		}

		ok, err := govalidator.ValidateStruct(employee)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}
