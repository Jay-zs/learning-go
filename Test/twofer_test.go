package Test

import "testing"

func TestTwofer(t *testing.T) {

	testcases := []struct {
		input          string
		expectedOutput string
	}{{

		input:          "",
		expectedOutput: "one for you , one for me",
	}, {

		input:          "Jay",
		expectedOutput: "one for Jay , one for me",
	}, {

		input:          "Shubham",
		expectedOutput: "one for Shubham , one for me",
	}, {

		input:          "Rohan",
		expectedOutput: "one for Rohan , one for me",
	},
	}

	for _, test := range testcases {

		res := Twofer(test.input)

		if res != test.expectedOutput {
			t.Error("test cases  fail")
		}

	}
}

func BenchmarkTwofer(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Twofer("Jay")
	}
}
