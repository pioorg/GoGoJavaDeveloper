package main

import (
	"fmt"
	"sync"
)

func printStr(s string) {
	for i := 0; i < 100; i++ {
		fmt.Println(s, i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		printStr("Hi")
	}()
	go func() {
		defer wg.Done()
		printStr("Hello")
	}()
	wg.Wait()
}
