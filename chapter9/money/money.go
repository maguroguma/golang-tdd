package money

type Money interface {
	GetAmount() int
	SetAmount(int)
	Equals(interface{}) bool
	Times(int) Money
	Currency() string
}

type money struct {
	amount   int
	currency string
}

func NewMoney(amount int, currency string) *money {
	return &money{amount, currency}
}

func (m *money) GetAmount() int {
	return m.amount
}
func (m *money) SetAmount(amount int) {
	m.amount = amount
}
func (m *money) Equals(comp_m *money) bool {
	return m.amount == comp_m.amount
}
func (m *money) GetCurrency() string {
	return m.currency
}

func MakeDollar(amount int) Money {
	return NewDollar(amount)
}
func MakeFranc(amount int) Money {
	return NewFranc(amount)
}
