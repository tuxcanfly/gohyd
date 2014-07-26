package main

import "fmt"

type PrettyString string

func (s PrettyString) String() string {
	return fmt.Sprintf("--'%#v'--", s)
}

func main() {
	ps := PrettyString("Hi I'm a string")
	fmt.Println(ps)
}
