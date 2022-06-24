package Test

import "strings"

func Twofer(name string) string {

	name = strings.Replace(name, " ", "", -1)
	name = strings.Replace(name, "\n", "", -1)
	var result string
	if name != "" {
		result = "one for " + name + " , one for me"
	} else {
		result = "one for " + "you" + " , one for me"
	}

	return result

}
