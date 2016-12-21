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

func calc(b benefits.Plan) {
	var youPay benefits.USD = 0.0
	var theyPay benefits.USD = 0.0
	for i := 0; i < 12; i++ {
		t, u := b.PayFor(foo.CostPerVisit)
		theyPay += t
		youPay += u
		fmt.Println(b)
		fmt.Printf("You've paid %v, They've paid %v\n", youPay, theyPay)
	}
}

func main() {
	calc(tweedleDee)
	calc(tweedleDumb)
}
