package benefits

type Percentage int
type USD float64
type Name string



// TODO(aoeu): Everything here considers out of network.

type Plan struct {
	Name
	Deductible         USD
	MaximumOutOfPocket USD
	Coinsurance        Percentage
}

type Coster interface {
	Cost() USD
}

type Freqer interface {
	Freq() int
}

type CosterFreqer interface {
	Coster
	Freqer
}

type Doctor struct {
	CostPerVisit  USD
	VisitsPerYear int
}

func (d Doctor) Cost() USD {
	return d.CostPerVisit
}

func (d Doctor) Freq() int {
	return d.VisitsPerYear
}

type Prescription struct {
	CostPerRefill  USD // Insured.
	RefillsPerYear int
}

func (p Prescription) Cost() USD {
	return p.CostPerRefill
}

func (p Prescription) Freq() int {
	return p.RefillsPerYear
}

func (b *Plan) PayFor(n USD) (theyPay USD, youPay USD) {
	remainder := b.subtractFromDeductible(n)
	switch {
	case b.Deductible > 0:
		if b.MaximumOutOfPocket > 0 {
			b.MaximumOutOfPocket -= n
		}
		return 0.0, n
	case b.Deductible == 0 && remainder == 0 && b.MaximumOutOfPocket > 0:
		if b.MaximumOutOfPocket > 0 {
			b.MaximumOutOfPocket -= n
		}
		return 0.0, n
	case b.MaximumOutOfPocket == 0:
		return n, 0.0
	}

	yourPercentage := float64(b.Coinsurance) * 0.01
	theirPercentage := 1.0 - yourPercentage

	theyPay = USD(float64(n) * theirPercentage)
	copay := USD(float64(n) * yourPercentage)

	youPay += copay

	if r := b.subtractFromMaximumOutOfPocket(copay); r > 0 {
		theyPay += r
		youPay -= r
	}

	return theyPay, youPay
}

func (b *Plan) Buy(c CosterFreqer) (theyPay USD, youPay USD) {
	for i := 0; i < c.Freq(); i++ {
		t, u := b.PayFor(c.Cost())
		theyPay += t
		youPay += u
	}
	return theyPay, youPay
}

func (b *Plan) subtractFromMaximumOutOfPocket(n USD) USD {
	b.MaximumOutOfPocket, n = decr(b.MaximumOutOfPocket, n)
	return n
}

func (b *Plan) subtractFromDeductible(n USD) USD {
	b.Deductible, n = decr(b.Deductible, n)
	return n
}

func decr(current USD, n USD) (next USD, leftover USD) {
	if current >= 0 {
		current -= n
		n = 0.0
	}
	switch {
	case current < 0:
		n += current * -1.0
		current = 0.00
	case current == 0:
		n = 0.0
	}
	return current, n
}

// Keep in mind that coinsurance only kicks in *after* deductible has been met.
