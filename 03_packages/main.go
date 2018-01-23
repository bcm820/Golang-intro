package main

import (
	"fmt"

	"github.com/bcmendoza/golang-basics/03_packages/packages"
)

func main() {
	fmt.Println(packages.Exported())
	fmt.Println("Hello", packages.Name)
}
