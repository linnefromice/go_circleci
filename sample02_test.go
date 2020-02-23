package sample

import (
	"testing"
)

func Test_plus(t *testing.T) {
	num1 := 8
	num2 := 4
	actual := 12
	expected := plus(num1, num2)
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func Test_minus(t *testing.T) {
	num1 := 8
	num2 := 4
	actual := 4
	expected := minus(num1, num2)
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func Test_multiple(t *testing.T) {
	num1 := 8
	num2 := 4
	actual := 32
	expected := multiple(num1, num2)
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func Test_divide(t *testing.T) {
	num1 := 8
	num2 := 4
	actual := 2
	expected := divide(num1, num2)
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
