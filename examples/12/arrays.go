package main

import "fmt"

func printA3(s3 [3]string) {
	fmt.Println(s3[0], s3[1], s3[2])
}

func main() {
	a3 := [3]string{"a", "b", "c"}
	printA3(a3)
	a3 = [...]string{"d", "e", "f"}
	printA3(a3)
	a3bis := a3
	a3bis[0] = "DD"
	printA3(a3)
	printA3(a3bis)

	// a3 = [...]string{"1", "2", "3", "4"}

	//a4 := [...]string{"1", "2", "3", "4"}
	//printA3(a4)
}
