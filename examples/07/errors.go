package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func fullName(uid int) (string, string, error) {
	if uid <= 0 {
		return "", "", errors.New("incorrect UID")
	}
	if uid == 1 {
		return "John", "Doe", nil
	}
	if uid == 2 {
		return "Susan", "Doe", nil
	}
	return "", "", errors.New("no such user")
}

func main() {
	uid := rand.Intn(3)
	fn, ln, err := fullName(uid)
	if err == nil {
		fmt.Println("Please welcome", fn, ln)
	} else {
		fmt.Println("User", uid, "unknown", err)
	}
}
