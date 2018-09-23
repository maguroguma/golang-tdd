package money

type Bank struct {
	rates map[Pair]int
}

func NewBank() *Bank {
	rates := make(map[Pair]int)
	return &Bank{rates}
}

func (b *Bank) Reduce(source Expression, to string) *Money {
	return source.Reduce(b, to)
}

func (b *Bank) AddRate(from, to string, rate int) {
	b.rates[Pair{from, to}] = rate
}

func (b *Bank) Rate(from, to string) int {
	if from == to {
		return 1
	}
	return b.rates[Pair{from, to}]
}