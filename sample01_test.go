package sample

import (
	"testing"
)

func Test_HelloWorld_01(t *testing.T) {
	actual := HelloWorld("hoge")
	expected := "Hello World, hoge!"
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func Test_HelloWorld_02(t *testing.T) {
	actual := HelloWorld("12345")
	expected := "Hello World, 12345!"
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func Test_HelloWorld_03(t *testing.T) {
	actual := HelloWorld("")
	expected := "Hello World!"
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
