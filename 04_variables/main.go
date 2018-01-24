package main

import "fmt"

// two primary ways to declare variables
// 1. shorthand: only used inside a func
// 2. using var to assign/initialize
// there are others, but these are preferred

func shorthand() {
	a := 10
	b := "golang"
	c := true
	fmt.Printf("%v - %T \n", a, a)
	fmt.Printf("%v - %T \n", b, b)
	fmt.Printf("%v - %T \n", c, c)
}

// %v prints the default value format
// %T prints the type

var d = "cowboy"
var e = "homeboy"

// you can declare multiple variables
// by listing them out!
// this also works using var

func multDeclare() {
	a, b, c, d := 1, false, "haha", `huh`
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

// var can be used to assign zero values
// these are usually done outside of func
// and in func they are assigned values

func zero() {
	var a int
	var b string
	var c float64
	var d bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func main() {
	shorthand()
	fmt.Printf("%v - %T \n", d, d)
	fmt.Printf("%v - %T \n", e, e)
	multDeclare()
	zero()
}
