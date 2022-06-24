package TestIntro

import (
	"testing"
)

type TestCases struct {
	desc   string
	input  int
	expOut bool
}

func TestPrime(t *testing.T) {
	testcases := []TestCases{
		{"prime", 2, true},
		{"non-prime", 4, false},
		{"edge case", 1, false},
	}
	for i, tc := range testcases {
		res, err := Prime(tc.input)
		if res != tc.expOut {
			t.Error(i+1, "Test Failed", err)
		}
	}

}
