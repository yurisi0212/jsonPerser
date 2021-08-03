package main

import "testing"

func TestNormalFormat(t *testing.T) {
	TestCase := "{ \"test:\"1\", }"
	output := checkJsonFormat(TestCase)
	if output != true {
		t.Errorf("Test Failed")
	}
}

func TestBrokenFormat(t *testing.T) {
	TestCase := "test:1"
	output := checkJsonFormat(TestCase)
	if output != false {
		t.Errorf("Test Failed")
	}
}
