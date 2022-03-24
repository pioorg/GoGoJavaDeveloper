package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Go!")
	fmt.Println("It's", time.Now().Format("2006-01-02T15:04:05.999999999 MST"))
}
