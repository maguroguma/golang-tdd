package constant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConst(t *testing.T) {
	const (
		INT_CONST = 100
		X
		Y
		STRING_CONST = "constant"
		S
		T
	)
	a := []int{INT_CONST, X, Y}
	aExpected := []int{100, 100, 100}
	b := []string{STRING_CONST, S, T}
	bExpected := []string{"constant", "constant", "constant"}
	assert.Exactly(t, aExpected, a)
	assert.Exactly(t, bExpected, b)
	assert.IsType(t, INT_CONST, 0)
	assert.IsType(t, STRING_CONST, "test string")
}
