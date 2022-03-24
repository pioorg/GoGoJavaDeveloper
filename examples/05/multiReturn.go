package main

import (
	"fmt"
)

func fullName() (string, string) {
	return "John", "Doe"
}

func main() {
	fn, ln := fullName()
	fmt.Println("Please welcome", fn, ln)
}
