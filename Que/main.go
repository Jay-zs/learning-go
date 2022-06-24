package main

import (
	"fmt"
	"strings"
)

//Returns the string with message "One for *name ,One for me"
func solve(name string) string {
	if strings.TrimSpace(name) == "" {
		return "One for you,One for me"
	}
	return "One for " + name + ",One for me"
}

func main() {
	fmt.Println(solve("bob"))
	fmt.Println(solve(" "))
	fmt.Println(solve("\n"))
}
