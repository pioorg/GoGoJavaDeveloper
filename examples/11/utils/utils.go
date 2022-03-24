package utils

import "fmt"

var Accessible = "this is accessible"
var inaccessible = "this is NOT accessible"

func SayHi() {
	fmt.Println("Hi!")
	fmt.Println(Accessible)
	fmt.Println(inaccessible)
}
