package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}

/*

2 main types of scope:
1. package scope
2. block scope

Closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope

In this example, x is scoped within the wrapper func, which when called,
returns an anonymous func that increments and returns x (which it has access
to since it is inside of wrapper's scope).

*/
