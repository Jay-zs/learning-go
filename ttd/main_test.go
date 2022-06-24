package ttd

import "testing"

func TestBob(t *testing.T) {
	testcases := []struct {
		desc   string
		input  string
		expOut string
	}{
		{"Question politely", "How are you?", "Sure."},
		{"Shouting", "HOW ARE YOU", "Whoa, chill out!"},
		{"Shouting Ques", "HOW ARE YOU?", "Calm down, I know what I'm doing!"},
		{"not addressing", "\n", "Fine. Be that way!"},
		{"anything else", "Hi", "Whatever."},
	}
	for _, tc := range testcases {
		if res := bob(tc.input); res != tc.expOut {
			t.Error("failed ", tc.expOut, res)
		}
	}
}
