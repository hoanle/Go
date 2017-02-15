package main

import (
	"testing"
)

type TestCase struct {
	input               string
	input_after_changed string
	expected            bool
}

func TestReplaceAllSpaces(t *testing.T) {
	tests := []TestCase{
		{"this is a test", "this%20is%20a%20test", true},
		{"this is a test", "this%20is%20a%20test%20", false},
		{"anothertest   ", "anothertest", true},
		{"anothertest   ", "anothertest%20", false},
	}

	for _, v := range tests {
		s := replaceAllSpaces(v.input)
		compare := (s == v.input_after_changed)
		if compare != v.expected {
			t.Fatalf("Input %s. after changed %s. Actual %s\n", v.input, v.input_after_changed, s)
		}
	}
}
