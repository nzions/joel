package main

import (
	"fmt"
	"joel/lib"
)

func main() {
	fmt.Println("Hello joe")

	c := &lib.Cisco{}

	lib.DoitWithCC(c)
	c.Foo()
}
