package money

type Dollar struct {
	amount int
}

func NewDollar(amount int) *Dollar {
	d := new(Dollar)
	d.amount = 10
	return d
}

func (d *Dollar) times(multiplier int) {
}
