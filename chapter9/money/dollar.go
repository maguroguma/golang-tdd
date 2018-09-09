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

func (d *Dollar) Times(multiplier int) Money {
	return NewDollar(d.money.GetAmount() * multiplier)
}
func (d *Dollar) GetAmount() int {
	return d.money.GetAmount()
}
func (d *Dollar) SetAmount(amount int) {
	d.money.SetAmount(amount)
}
func (d *Dollar) Equals(obj interface{}) bool {
	switch obj.(type) {
	case *Dollar:
		return d.money.Equals(obj.(*Dollar).money)
	case *Franc:
		return false
	default:
		return false
	}
}
