// testInterface
package main

import (
	"fmt"
)

type DataWriter interface {
	WriteData(data interface{}) error

	canWrite() bool
}

type file struct {
	abc int
}

func (d *file) WriteData(data interface{}) error {

	fmt.Println("WriteData:", data)
	return nil
}
func (d *file) canWrite() bool {
	return true
}

func getType(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("type is int")
	case string:
		fmt.Println("type is string")
	default:
		fmt.Println(123)
		fmt.Println("type is unknown")
	}
}

func main() {

	f := new(file)

	var writer DataWriter

	writer = f

	writer.WriteData("data")

	fmt.Println("Hello World!")
	xx := 100
	getType(xx)
	abc := "100"
	getType(abc)
	getType(f)
}
