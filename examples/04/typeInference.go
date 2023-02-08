package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// you don't see this often
	var rn1 int = rand.Intn(20)
	fmt.Println("We've got: ", rn1)

	// this is more frequent
	var rn2 = rand.Intn(20)
	fmt.Println("We've got: ", rn2)

	// most of the time you see this
	rn3 := rand.Intn(20)
	fmt.Println("We've got: ", rn3)
}
