// hello.go
package main

import (
	"fmt"
)

func main() {
	var num int
	num = 100
	var str string
	str = "abc"
	bb := true
	var money float64 = 56.78
	fmt.Println("Hello World!", num, str, bb, int(money))

	const size = 1001
	fmt.Println(size, "type of size is %T", bb)
}
