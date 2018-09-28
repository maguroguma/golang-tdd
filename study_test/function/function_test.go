package function

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOfAnonymousFunction(t *testing.T) {
	sum := func(x, y int) int {
		return x + y
	}

	assert.Equal(t, 10, sum(3, 7))
}

func TestOfMultiReturnFunction(t *testing.T) {
	div := func(first, second int) (int, int) {
		q := first / second
		r := first % second
		return q, r
	}

	q, r := div(22, 5)
	compared := []int{q, r}
	expected := []int{4, 2}
	assert.Exactly(t, expected, compared)
}

func TestOfFunctionType(t *testing.T) {
	sum := func(x, y int) int {
		return x + y
	}
	multi := func(x, y int) int {
		return x * y
	}
	div := func(first, second int) (int, int) {
		q := first / second
		r := first % second
		return q, r
	}

	assert.IsType(t, sum, multi)
	assert.IsType(t, sum, div) // fail
}
