package money

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	product := five.times(2)
	if 10 != product.amount {
		t.Error("同値エラー！")
	}
	product = five.times(3)
	if 15 != product.amount {
		t.Error("同値エラー！")
	}
}

func TestEquality(t *testing.T) {
	if !NewDollar(5).equals(NewDollar(5)) {
		t.Error("equalsエラー！")
	}
}
