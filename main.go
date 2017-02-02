package main

import "fmt"

func dummy(some int) int {
	return some + 5
}

func main() {
	fmt.Println(dummy(37))
}
