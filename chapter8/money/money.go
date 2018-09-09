package money

//type Money interface {
//	GetAmount() int
//	SetAmount(int) int
//	Equals(*Money) bool
//}

type Money struct {
	amount int
}

func (m *Money) GetAmount() int {
	return m.amount
}

func (m *Money) SetAmount(amount int) {
	m.amount = amount
}

func (m *Money) Equals(comp_m *Money) bool {
	return m.amount == comp_m.amount
}
