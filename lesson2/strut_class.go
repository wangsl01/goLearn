// strut_class
package main

import (
	"fmt"
)

type Base struct {
	name string
}

type sonBase struct {
	id   int
	base Base
}

func main() {
	var base Base = Base{"abc"}
	base1 := new(Base)
	fmt.Println(base.name)
	base1.name = "1111"
	fmt.Println(base1.name)
	//bb := 100
	fmt.Println(&base)

}
