package money

type Sum struct {
	augend *Money
	addend *Money
}

func NewSum(augend, addend *Money) *Sum {
	return &Sum{augend, addend}
}

func (s *Sum) Reduce(bank *Bank, to string) *Money {
	amount := s.augend.amount + s.addend.amount
	return NewMoney(amount, to)
}
