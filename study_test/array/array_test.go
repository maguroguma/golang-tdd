package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArray(t *testing.T) {
	a := [5]int{1, 2, 3}
	expected := [5]int{1, 2, 3, 0, 0}

	assert.Equal(t, expected, a)
	assert.Equal(t, 5, len(a))
}

func TestArrayDeepCopy(t *testing.T) {
	a := [3]int{1, 2, 3}
	b := a
	assert.Equal(t, a, b)

	b[2] = 100
	assert.NotEqual(t, a, b)
}
