package main

import (
	"fmt"
	"math/rand"
	"time"
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
	rand.Seed(time.Now().UnixNano())
	uid := rand.Intn(3)
	fn, ln, ok := fullName(uid)
	if ok {
		fmt.Println("Please welcome", fn, ln)
	} else {
		fmt.Println("User", uid, "unknown")
	}
}
