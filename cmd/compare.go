package main

import (
	"fmt"
	"github.com/aoeu/benefits"
)

var foo = benefits.Doctor{
	CostPerVisit:  200,
	VisitsPerYear: 12,
}

var bar = benefits.Prescription{
	CostPerRefill:  15.00,
	RefillsPerYear: 12,
}

var baz = benefits.Prescription{
	CostPerRefill:  35.00,
	RefillsPerYear: 12,
}

var tweedleDee = benefits.Plan{
	Name:               "Freedux Oxen Corp",
	Deductible:         750.00,
	MaximumOutOfPocket: 2750.00,
	Coinsurance:        20,
}

var tweedleDumb = benefits.Plan{
	Name:               "United Bullpuckey",
	Deductible:         5000.00,
	MaximumOutOfPocket: 10000.00,
	Coinsurance:        30,
}

func c(b benefits.Plan) (theyPay benefits.USD, youPay benefits.USD) {
	for i := 0; i < 12; i++ {
		t, u := b.PayFor(foo.Cost() * 4)
		theyPay += t
		youPay += u

		t, u = b.PayFor(bar.Cost())
		theyPay += t
		youPay += u

		t, u = b.PayFor(baz.Cost())
		theyPay += t
		youPay += u
	}
	return theyPay, youPay
}

func main() {
	t, u := c(tweedleDee)
	s := "On plan %v\n\tyou'd pay %v\n\tthey'd pay %v.\n"
	fmt.Printf(s, tweedleDee.Name, u, t)
	t, u = c(tweedleDumb)
	fmt.Printf(s, tweedleDumb.Name, u, t)
}
