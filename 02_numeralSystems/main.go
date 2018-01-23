package main

import "fmt"

/*
	Println prints on one line
	Printf formats according to specifier
*/

func main() {
	fmt.Printf("%d \n", 42) // %d prints as decimal
	fmt.Printf("%b \n", 42) // %b prints as binary
	fmt.Printf("%x \n", 42) // %x prints as hexidecimal
	fmt.Printf("%q \n", 42) // %x prints as UTF-8
	firstTwoHundred()
}

func firstTwoHundred() {
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
