package money

type Money struct {
	amount   int
	currency string
}

func NewMoney(amount int, currency string) *Money {
	return &Money{amount, currency}
}
func MakeDollar(amount int) *Money {
	return NewMoney(amount, "USD")
}
func MakeFranc(amount int) *Money {
	return NewMoney(amount, "CHF")
}

func (m *Money) Currency() string {
	return m.currency
}
func (m *Money) Times(multiplier int) *Money {
	return NewMoney(m.amount*multiplier, m.currency)
}
func (m *Money) Equals(comp_m *Money) bool {
	if m.amount == comp_m.amount && m.currency == comp_m.currency {
		return true
	} else {
		return false
	}
}
func (m *Money) Plus(addedMoney *Money) *Money {
	return NewMoney(m.amount+addedMoney.amount, m.currency)
}
