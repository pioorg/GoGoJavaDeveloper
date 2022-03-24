package main

import "fmt"

type Something struct {
}

func (s Something) do() {
	fmt.Println("nothing")
}
func (s Something) doAThing(thing string) {
	fmt.Println(thing)
}

func main() {
	s := &Something{}
	s.do()
	s.doAThing("something else")
}
