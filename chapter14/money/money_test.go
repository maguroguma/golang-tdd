package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	EqualAlert = "they should be equal!"
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

func TestSimpleAddition(t *testing.T) {
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

func TestPlusReturnsSum(t *testing.T) {
	var five *Money
	var result Expression
	var sum *Sum

	five = MakeDollar(5)
	result = five.Plus(five)
	sum = result.(*Sum)
	assert.Equal(t, five, sum.augend, EqualAlert)
	assert.Equal(t, five, sum.addend, EqualAlert)
}

func TestReduceSum(t *testing.T) {
	var sum Expression
	var bank *Bank
	var result *Money

	sum = NewSum(MakeDollar(3), MakeDollar(4))
	bank = NewBank()
	result = bank.Reduce(sum, "USD")
	assert.Equal(t, MakeDollar(7), result, EqualAlert)
}

// 回帰テスト
func TestIdentityRate(t *testing.T) {
	assert.Equal(t, 1, NewBank().Rate("USD", "USD"))
}

func TestReduceMoney(t *testing.T) {
	var bank *Bank
	var result *Money

	bank = NewBank()
	result = bank.Reduce(MakeDollar(1), "USD")
	assert.Equal(t, MakeDollar(1), result, EqualAlert)
}

func TestReduceMoneyDifferentCurrency(t *testing.T) {
	var result *Money
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	result = bank.Reduce(MakeFranc(2), "USD")
	assert.Equal(t, MakeDollar(1), result, EqualAlert)
}

// 学習用テスト
func TestArrayEquals(t *testing.T) {
	assert.Equal(t, []interface{}{"abc", "def"}, []interface{}{"abc", "def"}, EqualAlert)
	assert.ElementsMatch(t, []interface{}{"abc", "def"}, []interface{}{"abc", "def"}, EqualAlert)
}
