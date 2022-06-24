package calculator

import (
	"log"
	"strings"
)

func calculator(a, b int, op string) int {
	op = strings.TrimSpace(op)
	if op == "" {
		log.Println("Invalid Op")
		return -1
	}
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			return -1
		} else {
			return a / b
		}
	case "|":
		return a | b
	case "&":
		return a & b
	case "^":
		return a ^ b
	}
	return -1
}
