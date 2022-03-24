//go:build linux

package main

import (
	"fmt"
	"sync"
)

/*

#include <stdio.h>
#include <sys/types.h>
#include <sys/syscall.h>
#include <unistd.h>

static inline void greet(char* greeting) {
	pid_t tid = syscall(__NR_gettid);
	printf("%s from %d\n", greeting, tid);
}

*/
import "C"

func printStr(s string) {
	for i := 0; i < 100; i++ {
		greeting := C.CString(fmt.Sprintf("%s %d", s, i))
		C.greet(greeting)
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
