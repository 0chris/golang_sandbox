package main

/*
#cgo CFLAGS: -I./lib
#cgo LDFLAGS: -L./lib -Wl,-rpath=./lib -lshared
#include "lib/libshared.h"
*/
import "C"
import "fmt"

func main() {
	fmt.Println("Main")
	C.something()
}
