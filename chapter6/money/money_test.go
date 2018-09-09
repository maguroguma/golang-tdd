package money

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	product := five.Times(2)
	if !NewDollar(10).Equals(product) {
		t.Error("同値エラー！")
	}
	product = five.Times(3)
	if !NewDollar(15).Equals(product) {
		t.Error("同値エラー！")
	}
}

func TestEquality(t *testing.T) {
	if !NewDollar(5).Equals(NewDollar(5)) {
		t.Error("Dollar Equalsエラー！")
	}
	if NewDollar(5).Equals(NewDollar(6)) {
		t.Error("Dollar Equalsエラー！")
	}

	if !NewFranc(5).Equals(NewFranc(5)) {
		t.Error("Franc Equalsエラー！")
	}
	if NewFranc(5).Equals(NewFranc(6)) {
		t.Error("Franc Equalsエラー！")
	}
}

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	product := five.Times(2)
	if !NewFranc(10).Equals(product) {
		t.Error("同値エラー！")
	}
	product = five.Times(3)
	if !NewFranc(15).Equals(product) {
		t.Error("同値エラー！")
	}
}
