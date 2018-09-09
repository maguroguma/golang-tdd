package money

type Franc struct {
	*money
}

func NewFranc(amount int) *Franc {
	f := new(Franc)
	f.money = new(money)
	f.money.setAmount(amount)
	return f
}

func (f *Franc) Times(multiplier int) *Franc {
	return NewFranc(f.money.getAmount() * multiplier)
}

func (f *Franc) GetAmount() int {
	return f.money.getAmount()
}

func (f *Franc) SetAmount(amount int) {
	f.money.setAmount(amount)
}

func (f *Franc) Equals(comp_f *Franc) bool {
	return f.money.equals(comp_f.money)
}
