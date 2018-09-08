package money

type Dollar struct {
	amount int
}

type Franc struct {
	amount int
}

func NewDollar(amount int) *Dollar {
	d := new(Dollar)
	d.amount = amount
	return d
}

func NewFranc(amount int) *Franc {
	f := new(Franc)
	f.amount = amount
	return f
}

func (d *Dollar) times(multiplier int) *Dollar {
	return NewDollar(d.amount * multiplier)
}

func (d *Dollar) equals(comp_d *Dollar) bool {
	return d.amount == comp_d.amount
}

func (f *Franc) times(multiplier int) *Franc {
	return NewFranc(f.amount * multiplier)
}

func (f *Franc) equals(comp_f *Franc) bool {
	return f.amount == comp_f.amount
}
