// hello
package main

import (
	"fmt"
)

var bbbb = 1100

func old_main() {
	fmt.Println("Hello World!")
	var abc string = "abc"
	intArg := 1
	fmt.Println(abc, intArg)
	bc, ef := "111", 111
	fmt.Println(bc, ef)
	var bb bool
	fmt.Println(bb)
	fmt.Println(string(intArg), "=====")
	cc := &bc
	fmt.Println(&bc, *cc)
	fmt.Println(bbbb + 10)
	const (
		ip   = "192.168.1.1"
		port = 1234
	)
	fmt.Println(ip)

	type Weekday int
	const (
		Sunday Weekday = iota + 1
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Saturday)

	var intArray [10]int
	intArray[0] = 0
	fmt.Println(intArray[0])
	fmt.Println(intArray[len(intArray)-1])

	fmtArray := [...]string{"1A", "2B", "3C"}
	fmt.Println(fmtArray[0])
	fmt.Println(len(fmtArray))

	for k, v := range fmtArray {
		fmt.Println(k, v)
	}

	slice1 := fmtArray[0:]
	fmt.Println(slice1)
	slice1 = slice1[0:2]
	fmt.Println(slice1)
	slice1 = slice1[:1]
	fmt.Println(slice1)
	slice1 = fmtArray[:]
	fmt.Println(slice1)

	var intList []int
	fmt.Println(intList)
	intList = make([]int, 2, 10)
	fmt.Println(intList)
	intList = append(intList, 3)
	fmt.Println(intList)
	intList = append(intList, []int{10, 11, 12}...)
	fmt.Println(intList)

	intRet, strRet := testFor(9, 9)
	fmt.Println(intRet, strRet)
	int1, int2 := testParamName(110, 299)
	fmt.Println(int1, int2)

	var funcProxy func(int, int) (int, string)
	funcProxy = testFor
	funcProxy(19, 19)

	nf := func(name string) string {
		return "Hello World:" + name
	}
	ret := nf("wangsl")
	fmt.Println(ret)

	var aax Integer = 100
	var xx, yy = aax.less(2, "aaa")
	fmt.Println(xx, yy)

	aax.Add(1001)
	fmt.Println(aax)
}

/**
* 变更入参数据 += 传*ss
* 方法内赋值   = 传ss
 */
func (ss *Integer) Add(b Integer) {
	*ss += b

}

type Integer int

func (aa1 Integer) less(bb1 Integer, str string) (bool, string) {
	fmt.Println(aa1)
	return aa1 < bb1, str
}

func testFor(x, y int) (int, string) {
	if x < 1 || y < 1 {
		return 0, "empty"
	}
	for i := 0; i <= x; i++ {
		for k := 0; k <= i; k++ {
			fmt.Printf("%d*%d=%d ", i, k, i*k)
		}
		fmt.Println()
	}
	return 200, "success"
}

func testParamName(x, y int) (a, b int) {
	if x < 1 || y < 1 {
		return 0, 0
	} else {
		a = x
		b = y
		return a, b
	}
}
