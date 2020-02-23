package sample

import (
	"testing"
)

func Test_plus_01(t *testing.T) {
	num1 := 8
	num2 := 4
	actual := 12
	expected := plus(num1, num2)
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func Test_plus_02(t *testing.T) {
	num1 := 0
	num2 := -1
	actual := -1
	expected := plus(num1, num2)
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func Test_minus_01(t *testing.T) {
	num1 := 8
	num2 := 4
	actual := 4
	expected := minus(num1, num2)
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func Test_minus_02(t *testing.T) {
	num1 := 0
	num2 := -1
	actual := 1
	expected := minus(num1, num2)
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func Test_multiple_01(t *testing.T) {
	num1 := 8
	num2 := 4
	actual := 32
	expected := multiple(num1, num2)
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func Test_multiple_02(t *testing.T) {
	num1 := 0
	num2 := -1
	actual := 0
	expected := multiple(num1, num2)
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func Test_multiple_03(t *testing.T) {
	num1 := -5
	num2 := -7
	actual := 35
	expected := multiple(num1, num2)
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func Test_divide_01(t *testing.T) {
	num1 := 8
	num2 := 4
	actual := 2
	expected := divide(num1, num2)
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func Test_divide_02(t *testing.T) {
	num1 := 0
	num2 := 4
	actual := 0
	expected := divide(num1, num2)
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func Test_divide_03(t *testing.T) {
	num1 := -8
	num2 := -8
	actual := 1
	expected := divide(num1, num2)
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
