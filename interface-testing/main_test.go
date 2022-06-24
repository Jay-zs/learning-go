package main

import "testing"

func TestValidate(t *testing.T) {
	tc := struct {
		desc         string
		input        employee
		expectedAged bool
		expectedEmp  employee
	}{
		"emp having age less than 22",
		employee{20, "tushar", "hsr"},
		false,
		employee{0, "", ""},
	}
	resAged, resEmp := validate(tc.input)
	if resAged != tc.expectedAged || resEmp != tc.expectedEmp {
		t.Error("failed")
	}
}
