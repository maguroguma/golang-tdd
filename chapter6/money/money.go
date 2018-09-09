package money

//type Money interface {
//	GetAmount() int
//	SetAmount(int) int
//	Equals(*Money) bool
//}

type money struct {
	amount int
}

func (m *money) getAmount() int {
	return m.amount
}

func (m *money) setAmount(amount int) {
	m.amount = amount
}

func (m *money) equals(comp_m *money) bool {
	return m.amount == comp_m.amount
}
