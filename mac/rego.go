// rego
package main

import (
	"fmt"
)

func test(x int) {
	defer println("a")
	defer println("b")

	defer func() {
		println(100 / x) // div0 异常未被捕获，逐步往外传递，最终终止进程。
	}()

	defer println("c")
}

func main1() {
	test(1)
	fmt.Println("Hello World!")
}

//func main() {
//	fmt.Println("Hello World!")
//}
