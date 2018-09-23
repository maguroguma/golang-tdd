package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	var five, product *Money
	five = MakeDollar(5)
	product = five.Times(2)
	assert.Equal(t, MakeDollar(10), product, "they should be equal!")
	product = five.Times(3)
	assert.Equal(t, MakeDollar(15), product, "they should be equal!")
}

func TestEquality(t *testing.T) {
	assert.True(t, MakeDollar(5).Equals(MakeDollar(5)), "Dollar Equal!")
	assert.False(t, MakeDollar(5).Equals(MakeDollar(6)), "Dollar Not Equal!")
	assert.False(t, MakeFranc(5).Equals(MakeDollar(5)), "Franc and Dollar Not Equal!")
}

func TestCurrency(t *testing.T) {
	assert.Equal(t, "USD", MakeDollar(1).Currency(), "they should be same currency!")
	assert.Equal(t, "CHF", MakeFranc(1).Currency(), "they should be same currency!")
}

func TestSimpleAddition(t * testing.T) {
	var five *Money
	var sum Expression
	var bank *Bank
	var reduced *Money

	five = MakeDollar(5)
	sum = five.Plus(five)
	bank = NewBank()
	reduced = bank.Reduce(sum, "USD")
	assert.Equal(t, MakeDollar(10), reduced, "they should be equal!")
}
