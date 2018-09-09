package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	product := five.Times(2)
	assert.Equal(t, NewDollar(10), product, "they should be equal!")
	product = five.Times(3)
	assert.Equal(t, NewDollar(15), product, "they should be equal!")
}

func TestEquality(t *testing.T) {
	assert.True(t, NewDollar(5).Equals(NewDollar(5)), "Dollar Equal!")
	assert.False(t, NewDollar(5).Equals(NewDollar(6)), "Dollar Not Equal!")
	assert.True(t, NewFranc(5).Equals(NewFranc(5)), "Franc Equal!")
	assert.False(t, NewFranc(5).Equals(NewFranc(6)), "Franc Not Equal!")
	assert.False(t, NewFranc(5).Equals(NewDollar(5)), "Franc and Dollar Not Equal!")
	assert.False(t, NewDollar(5).Equals(NewFranc(5)), "Dollar and Franc Not Equal!")
}

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	product := five.Times(2)
	assert.Equal(t, NewFranc(10), product, "they should be equal!")
	product = five.Times(3)
	assert.Equal(t, NewFranc(15), product, "they should be equal!")
}
