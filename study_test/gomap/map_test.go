package gomap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapBasic(t *testing.T) {
	m := map[int]string{
		1: "first",
		2: "second",
		3: "third",
	}
	_, ok1 := m[1]
	_, ok2 := m[0]
	assert.True(t, ok1)
	assert.False(t, ok2)
}

func TestMapDelete(t *testing.T) {
	m := map[int]string{
		1: "first",
		2: "second",
		3: "third",
	}
	delete(m, 1)
	_, ok := m[1]
	assert.False(t, ok)
}
