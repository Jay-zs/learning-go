package calculator

import "testing"

type input struct {
	a, b int
	op   string
}

func TestCalculator(t *testing.T) {
	testcases := []struct {
		desc   string
		inp    input
		ExpOut int
	}{
		{"add", input{1, 2, "+"}, 3},
		{"subtract", input{2, 1, "-"}, 1},
		{"multiply", input{1, 2, "*"}, 2},
		{"divide", input{2, 2, "/"}, 1},
		{"or op", input{1, 0, "|"}, 1},
		{"and op", input{1, 0, "&"}, 0},
		{"xor op", input{1, 0, "^"}, 1},
		{"empty operator", input{1, 2, ""}, -1},
		{"invalid operator", input{1, 2, "!"}, -1},
		{"op with spaces", input{1, 2, "     +"}, 3},
		{"divide by zero", input{2, 0, "/"}, -1},
	}
	for i, tc := range testcases {
		res := calculator(tc.inp.a, tc.inp.b, tc.inp.op)
		if res != tc.ExpOut {
			t.Error(i, "failed")
		}
	}
}
