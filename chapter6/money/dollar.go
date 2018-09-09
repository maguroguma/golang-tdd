package money

type Dollar struct {
	*money
}

func NewDollar(amount int) *Dollar {
	d := new(Dollar)
	d.money = new(money)
	d.money.setAmount(amount)
	return d
}

func (d *Dollar) Times(multiplier int) *Dollar {
	return NewDollar(d.money.getAmount() * multiplier)
}

func (d *Dollar) GetAmount() int {
	return d.money.getAmount()
}

func (d *Dollar) SetAmount(amount int) {
	d.money.setAmount(amount)
}

func (d *Dollar) Equals(comp_d *Dollar) bool {
	return d.money.equals(comp_d.money)
}
