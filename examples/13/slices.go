package main

import "fmt"

func printSlice(s []string) {
	for _, e := range s {
		fmt.Print(e, " ")
	}
	fmt.Println()
}

func main() {
	s := []string{"a", "b", "c"}
	printSlice(s)
	s = []string{"d", "e", "f"}
	printSlice(s)
	fmt.Println("How will the replacing work?")
	sBis := s
	sBis[0] = "DD"
	printSlice(s)
	printSlice(sBis)
	fmt.Println("How will the appending work?")
	sBis = append(sBis, "G")
	printSlice(sBis)
	fmt.Println("How will the replacing work AGAIN?")
	sBis[0] = "d"
	printSlice(s)
	printSlice(sBis)
}
