package collatz

import (
	"reflect"
	"testing"
)

func TestCollatz(t *testing.T) {
	hailstone, err := Collatz(10)
	if err != nil {
		t.Error(err)
	}
	want := []int{10, 5, 16, 8, 4, 2, 1}
	if reflect.DeepEqual(hailstone, want) == false {
		t.Errorf("want %v, got %v", want, hailstone)
	}
}

func TestEven(t *testing.T) {
	got, err := Even(10)
	if err != nil {
		t.Error(err)
	}
	want := 5
	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestOdd(t *testing.T) {
	got, err := Odd(10)
	if err != nil {
		t.Error(err)
	}
	want := 31
	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}
