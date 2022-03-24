package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type IncorrectUid struct {
	uid int
}

func (e IncorrectUid) Error() string {
	return fmt.Sprintf("incorrect user is %d", e.uid)
}

func fullName(uid int) (string, string, error) {
	if uid == 0 {
		return "", "", IncorrectUid{uid}
	}
	if uid == 1 {
		return "John", "Doe", nil
	}
	if uid == 2 {
		return "Susan", "Doe", nil
	}
	return "", "", errors.New("no such user")
}

func printError(err error) {
	u := &IncorrectUid{}
	if errors.As(err, u) {
		fmt.Printf(`incorrect UID: "%d"`, u.uid)
	} else {
		fmt.Printf("generic error: %v", err)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	uid := rand.Intn(3)
	fn, ln, err := fullName(uid)
	if err == nil {
		fmt.Println("Please welcome", fn, ln)
	} else {
		printError(err)
	}
}
