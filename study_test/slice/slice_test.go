package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicOperation(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, []int{2, 3, 4}, s[1:4])

	u := s[1:4]
	assert.Equal(t, []int{2, 3, 4, 5, 6, 7}, append(u, 5, 6, 7))
}

func TestSliceParameter(t *testing.T) {
	sum := func(params ...int) int {
		n := 0
		for _, v := range params {
			n += v
		}
		return n
	}
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, 55, sum(s...))
}
