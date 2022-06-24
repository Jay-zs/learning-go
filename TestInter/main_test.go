package TestInter

import "testing"

func TestArea(t *testing.T) {
	testcases := []struct {
		desc   string
		shape  geometry
		expOut float64
	}{
		{"rectangle", rectangle{2, 2}, 4},
		{"triangle", triangle{0, 0, 0, 2, 2}, 2},
		{"circle", circle{7}, 153.93},
		{"square", square{2}, 4},
		{"shape with negative values", square{-2}, -1},
	}
	for i, tc := range testcases {
		res := tc.shape.area()
		if res != tc.expOut {
			t.Error(i, "failed")
		}
	}
}

func TestPerimeter(t *testing.T) {
	testcases := []struct {
		desc   string
		shape  geometry
		expOut float64
	}{
		{"rectangle", rectangle{2, 2}, 8},
		{"triangle", triangle{2, 2, 2, 0, 0}, 6},
		{"circle", circle{7}, 43.98},
		{"square", square{2}, 8},
		{"shape with negative values", square{-2}, -1},
	}
	for i, tc := range testcases {
		res := tc.shape.perimeter()
		if res != tc.expOut {
			t.Error(i, "failed")
		}
	}
}
