package money

type Franc struct {
	*money
}

func NewFranc(amount int) *Franc {
	f := new(Franc)
	f.money = new(money)
	f.money.SetAmount(amount)
	return f
}

func (f *Franc) Times(multiplier int) *Franc {
	return NewFranc(f.money.GetAmount() * multiplier)
}

func (f *Franc) GetAmount() int {
	return f.money.GetAmount()
}

func (f *Franc) SetAmount(amount int) {
	f.money.SetAmount(amount)
}

func (f *Franc) Equals(comp_f *Franc) bool {
	return f.money.Equals(comp_f.money)
}
