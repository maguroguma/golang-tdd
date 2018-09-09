package money

type Franc struct {
	*money
}

func NewFranc(amount int) *Franc {
	f := new(Franc)
	f.money = NewMoney(amount, "CHF")
	return f
}

func (f *Franc) Times(multiplier int) Money {
	return NewFranc(f.money.GetAmount() * multiplier)
}
func (f *Franc) GetAmount() int {
	return f.money.GetAmount()
}
func (f *Franc) SetAmount(amount int) {
	f.money.SetAmount(amount)
}
func (f *Franc) Equals(obj interface{}) bool {
	switch obj.(type) {
	case *Dollar:
		return false
	case *Franc:
		return f.money.Equals(obj.(*Franc).money)
	default:
		return false
	}
}

func (f *Franc) Currency() string {
	return f.money.GetCurrency()
}
