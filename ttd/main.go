package ttd

import (
	"regexp"
	"strings"
)

func bob(inp string) string {
	if inp[len(inp)-1] == '?' {
		if str := strings.TrimSpace(inp); str == "?" {
			return "Fine. Be that way!"
		}
		if strings.ToUpper(inp) == inp {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	} else {
		isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
		if strings.IndexFunc(inp, isNotDigit) == 0 {
			if str := strings.TrimSpace(inp); str == "" {
				return "Fine. Be that way!"
			}
			if strings.ToUpper(inp) == inp {
				return "Whoa, chill out!"
			}
		} else {
			reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
			noNum := reg.ReplaceAllString(inp, "")
			if strings.ToUpper(noNum) == noNum {
				return "Whoa, chill out!"
			}
		}
		return "Whatever."
	}
}
