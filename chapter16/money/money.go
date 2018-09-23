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
func (m *Money) Times(multiplier int) Expression {
	return NewMoney(m.amount*multiplier, m.currency)
}
func (m *Money) Equals(comp_m *Money) bool {
	if m.amount == comp_m.amount && m.currency == comp_m.currency {
		return true
	} else {
		return false
	}
}
func (m *Money) Plus(addend Expression) Expression {
	return NewSum(m, addend)
}
func (m *Money) Reduce(bank *Bank, to string) *Money {
	rate := bank.Rate(m.currency, to)
	return NewMoney(m.amount / rate, to)
}