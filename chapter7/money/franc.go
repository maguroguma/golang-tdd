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

func (f *Franc) Equals(obj interface{}) bool {
	switch obj.(type) {
	case *Dollar:
		return f.money.equals(obj.(*Dollar).money)
	case *Franc:
		return f.money.equals(obj.(*Franc).money)
	default:
		return false
	}
}
