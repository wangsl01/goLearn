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
