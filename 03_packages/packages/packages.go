package packages // non-main packages match folder name

func notExported() string {
	return "Lower-case functions are not exported"
}

// Exported is an exported function
// All exported functions have comments
func Exported() string {
	return notExported() + ", but upper-case functions are exported!"
}

// Name is an exported variable
var Name = "Brian"

/*

Everything within this package is accessible from every .go file in the same package, because they share package scope. In order to access its exported (visible) variables/functions from outside the package, it must be imported.

The ONLY exception to this is the main package. You can't put everything in your main package and access everything across files. One main file with main func. Or, when you run your program, call "go run *.go" and it will run everything.

*/
