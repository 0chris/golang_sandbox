package main

import "fmt"
import "C"

//export something
func something() {
	fmt.Println("I'm in")
}

func main() {}
