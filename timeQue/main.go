package main

import (
	"fmt"
	"math"
	"time"
)

func afterGiga(input time.Time) time.Time {
	res := input.Add(time.Duration(math.Pow10(9)) * time.Second)
	return res
}
func main() {
	t, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	fmt.Println(t)
	fmt.Println(afterGiga(t))
}
