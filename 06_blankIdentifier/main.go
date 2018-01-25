package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	a := "a"
	b := "b" // error! b is declared and not used
	fmt.Println(a)
}

// In Go, everything declared must be used. No pollution!

func getReq() {
	res, err := http.Get("http://www.google.com")
	if err != nil {
		log.Fatal(err)
	}
	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", page)
}

// But you can tell a compiler you aren't using a var
// If you want, you can use a blank identifier...

func noErrHandling() {
	res, _ := http.Get("http://www.google.com")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
