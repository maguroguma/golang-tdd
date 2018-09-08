package money

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	five.times(2)
	if 10 != five.amount {
		t.Error("同値エラー！")
	}
	five.times(3)
	if 15 != five.amount {
		t.Error("同値エラー！")
	}
}
