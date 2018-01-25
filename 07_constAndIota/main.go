package main

import "fmt"

const p = "death & taxes"

func main() {
	const q = 42
	fmt.Println(a, b, c, d, e, f)
}

/*
constants are unchanging values
they don't need to be used
they can also be multi-initialized
*/

const (
	pi   = 3.14
	lang = "Go"
)

/*
an iota is an incrementing number
it has many interesting characteristics
*/

const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
)

const (
	d = iota // 0
	e        // 1
	f        // 2
)
