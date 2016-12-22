package benefits

import (
	"testing"
)

func TestDeductibleLowers(t *testing.T) {
	n := 7
	p := &Plan{Deductible: USD(n)}
	for i := 1; i < n; i++ {
		tt, u := p.PayFor(1)
		switch {
		case u != 1:
			t.Errorf("Expected you to pay 1, but was %v, deductible", u, p.Deductible)
		case tt != 0:
			t.Errorf("Expected them to pay 0, but was %v", tt)
		case p.Deductible != USD(n-i):
			t.Errorf("Expected deductible to decrease, but it did not %v", p.Deductible)
		}
	}
	tt, u := p.PayFor(1)
	switch {
	case u != 0:
		t.Errorf("Expected you to pay 0,  but was %v", u)
	case tt != 1:
		t.Errorf("Expected them to pay 1, but was %v", tt)
	}
}

func TestAllTheThings(t *testing.T) {
	p := &Plan{
		Deductible:         USD(100),
		MaximumOutOfPocket: USD(200),
		Coinsurance:        10,
	}

	// Drain the deductible.
	for i := 1; i <= 4; i++ {
		tt, u := p.PayFor(25)
		switch {
		case u != 25:
			t.Errorf("wrong %v, %v", u, p.Deductible)
		case tt != 0:
			t.Errorf("wrong")
		case p.Deductible != USD(100-(25*i)):
			t.Errorf("wrong %v", p.Deductible)
		}
	}
	if p.Deductible != 0 {
		t.Errorf("Expected all deductible paid, but there remains %v", p.Deductible)
	}
	if p.MaximumOutOfPocket != 100 {
		t.Errorf("Expected all out-of-pocket maximum to be equail to itself minus the deductible, but it is actually %v", p.MaximumOutOfPocket)
	}
	// Drain the max-out-of-pocket
	for i := 1; i <= 40; i++ {
		tt, u := p.PayFor(25)
		switch {
		case p.MaximumOutOfPocket != USD(100.0-(2.5*float64(i))):
			t.Errorf("%v %v", p.MaximumOutOfPocket, USD(100.0-(2.5*float64(i))))
		case tt != USD(25*.9):
			t.Errorf("")
		case u != USD(25*.1):
			t.Errorf("%v %v", u, USD(2.5))
		}
	}
	if p.MaximumOutOfPocket != 0 {
		t.Errorf("Expected all out-of-pocket maximum to be met, but there remains %v", p.MaximumOutOfPocket)
	}

	for i := 1; i <= 5; i++ {
		tt, u := p.PayFor(USD(10 * i))
		switch {
		case u != 0:
			t.Errorf("")
		case tt != USD(10*i):
			t.Errorf("")
		}
	}
}
