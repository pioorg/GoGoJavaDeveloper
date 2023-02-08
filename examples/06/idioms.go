package main

import (
	"fmt"
	"math/rand"
)

func fullName(uid int) (string, string, bool) {
	if uid == 1 {
		return "John", "Doe", true
	}
	if uid == 2 {
		return "Susan", "Doe", true
	}
	return "", "", false
}

func main() {
	uid := rand.Intn(3)
	fn, ln, ok := fullName(uid)
	if ok {
		fmt.Println("Please welcome", fn, ln)
	} else {
		fmt.Println("User", uid, "unknown")
	}
}
