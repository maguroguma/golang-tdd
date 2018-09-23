package money

type Sum struct {
	augend, addend Expression
}

func NewSum(augend, addend Expression) *Sum {
	return &Sum{augend, addend}
}

func (s *Sum) Reduce(bank *Bank, to string) *Money {
	//amount := s.augend.amount + s.addend.amount
	amount := s.augend.Reduce(bank, to).amount + s.addend.Reduce(bank, to).amount
	return NewMoney(amount, to)
}

func (s *Sum) Plus(addend Expression) Expression {
	return nil
}