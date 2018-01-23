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
