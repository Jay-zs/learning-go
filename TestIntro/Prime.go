package TestIntro

import (
	"math"
)

func Prime(num int) (bool, string) {
	if num < 2 {
		return false, "Neither prime nor composite"
	}
	sqRoot := int(math.Sqrt(float64(num)))
	for i := 2; i <= sqRoot; i++ {
		if num%i == 0 {
			return false, ""
		}
	}
	return true, ""
}
