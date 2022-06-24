package employee

import (
	"reflect"
	"testing"
)

func TestEmp(t *testing.T) {
	TestCases := []struct {
		desc      string
		inputEmp  employee
		ExpOutEmp employee
	}{
		{"1 yr exp", employee{"Jay", "11/03/2002", 1, 0, 0}, employee{"Jay", "11/03/2002", 1, 20, 20000}},
		{"2 to 5 yr exp", employee{"Jay", "11/03/2001", 3, 0, 0}, employee{"Jay", "11/03/2001", 3, 21, 50000}},
		{"5 to 10 yr exp", employee{"Jay", "11/03/2002", 7, 0, 0}, employee{"Jay", "11/03/2002", 7, 20, 100000}},
		{"more than 10 yr exp", employee{"Jay", "11/03/2002", 11, 0, 0}, employee{"Jay", "11/03/2002", 11, 20, 200000}},
		{"yoe is -ve", employee{"Jay", "11/03/2002", -1, 0, 0}, employee{"Jay", "11/03/2002", -1, 20, 0}},
		{"dob not provided", employee{"Jay", "", 11, 0, 0}, employee{"Jay", "", 11, 0, 200000}},
		{"yoe is -ve and dob not provided", employee{"Jay", "", -1, 0, 0}, employee{"Jay", "", -1, 0, 0}},
		{"dob not provided", employee{"Jay", "\n", 0, 0, 0}, employee{"Jay", "\n", 0, 0, 0}},
		{"name not provided", employee{"", "11/03/2002", 0, 0, 0}, employee{"", "11/03/2002", 0, 20, 0}},
	}
	for i, tc := range TestCases {
		OutEmp := emp(tc.inputEmp)
		if !reflect.DeepEqual(OutEmp, tc.ExpOutEmp) {
			t.Error("failed", i)
		}
	}
}
