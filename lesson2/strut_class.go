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
	base *Base
}

func (sonbase *sonBase) showId(id1 int) int {
	fmt.Println(sonbase.base.name)
	return sonbase.id + id1
}

func main() {
	var base Base = Base{"abc"}
	base1 := new(Base)
	fmt.Println(base.name)
	base1.name = "1111"
	fmt.Println(base1.name)
	sonbase := &sonBase{100, base1}
	fmt.Println(sonbase)
	newId := sonbase.showId(101)
	fmt.Println(newId)
}
