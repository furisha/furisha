package main

import (
	"fmt"
	)

func main() {
	fmt.Println("Hello everyone, im the best go programmer in the world")

	foo()

	fmt.Println("sdssdsgs")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I'm in foo")
}

func bar()  {
	fmt.Println("Ludnica")
}

// control flow:
// (1) sequence
// (2) loop; iterative
// (3) conditional
