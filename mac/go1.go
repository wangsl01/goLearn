// go1
package main

import (
	"bytes"
	"fmt"
	"math"
	"strings"
)

//bool
//string
//int、int8、int16、int32、int64
//uint、uint8、uint16、uint32、uint64、uintptr
//byte // uint8 的别名
//rune // int32 的别名 代表一个 Unicode 码
//float32、float64
//complex64、complex128

var name string = "wangsl"

/**
  jiandangongju
*/
func mainssss() {
	var abcInt int
	var abcInt8 int8 = 10
	var abcBool bool
	var abcFloat32 float32
	var abcInt64 = 1000000
	fmt.Println(abcInt, abcInt8, abcBool, abcFloat32, abcInt64)
	fmt.Println("=========================")
	bb := 10
	aa, bb := 100, 1000
	aa, bb = bb, aa
	fmt.Println(aa, bb)
	aa, _ = bb, aa
	fmt.Println(aa, bb)
	name = "wslzhq"
	fmt.Println(name)
	//形式参数会作为函数的局部变量来使用。
	fmt.Println("=========================")
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)
	fmt.Println("=========================")
	for i := 0; i < len(name); i++ {
		fmt.Println(i, name[i])
	}
	for c := range name {
		fmt.Printf("%d %c\n", c, name[c])
	}
	fmt.Println("=========================")
	pos := strings.Index(name, "h")
	lpos := strings.LastIndex(name, "s")
	fmt.Println(pos, lpos)
	fmt.Println("=========================")
	var stringBuilder bytes.Buffer
	stringBuilder.WriteString(name)
	stringBuilder.WriteString("home")
	fmt.Println(stringBuilder.String())
	fmt.Println("=========================")
	/* meiju */
	const (
		//changdu
		length = iota
		width
		height
	)
	fmt.Println(length, width, height)
	fmt.Println("=========================")
	var intArr [3]int
	for i := 0; i < 3; i++ {
		intArr[i] = i + 1
	}
	for k, v := range intArr {
		fmt.Println(k, v)
	}
	fmt.Println("======================")
	var slinceTest []string
	fmt.Println(slinceTest)
	fmt.Println("======================")
	var arr = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr[0], arr[1] = arr[1], arr[0]
	fmt.Println(arr)
	fmt.Println("======================")
	fmt.Println(resolveTime(1000000))

}

const (
	// 定义每分钟的秒数
	SecondsPerMinute = 60
	// 定义每小时的秒数
	SecondsPerHour = SecondsPerMinute * 60
	// 定义每天的秒数
	SecondsPerDay = SecondsPerHour * 24
)

func resolveTime(second int) (day, hour, minute int) {
	day = second / SecondsPerDay
	hour = second / SecondsPerHour
	minute = second / SecondsPerMinute
	return
}
