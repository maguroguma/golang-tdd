package money

type Dollar struct {
	*money
}

func NewDollar(amount int) *Dollar {
	d := new(Dollar)
	d.money = new(money)
	d.money.SetAmount(amount)
	return d
}

func (d *Dollar) Times(multiplier int) *Dollar {
	return NewDollar(d.money.GetAmount() * multiplier)
}

func (d *Dollar) GetAmount() int {
	return d.money.GetAmount()
}

func (d *Dollar) SetAmount(amount int) {
	d.money.SetAmount(amount)
}

func (d *Dollar) Equals(comp_d *Dollar) bool {
	return d.money.Equals(comp_d.money)
}
