package main

import (
	"fmt"
)

type Swimmer struct {
	style string
}

func (s Swimmer) Swim() {
	fmt.Println("swimming", s.style)
}

type Runner struct {
	style string
}

func (r Runner) Run() {
	fmt.Println("running", r.style)
}

type Hippo struct {
	Swimmer
	Runner
}

func main() {
	h := Hippo{
		Swimmer{"slow"},
		Runner{"fast"},
	}
	// this is pretty obvious
	h.Swimmer.Swim()
	h.Runner.Run()
	// how about this?
	h.Swim()
	h.Run()
}
