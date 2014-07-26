package main

import "fmt"

func main() {
	for i := 1; i < 20; i++ {
		s := ""
		if i%3 == 0 {
			s += "Fizz"
		}
		if i%5 == 0 {
			s += "Buzz"
		}
		if len(s) == 0 {
			fmt.Println(i)
		} else {
			fmt.Println(s)
		}
	}
}
