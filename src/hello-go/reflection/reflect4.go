package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num = 1.2345
	fmt.Println("old value of pointer: ", num)

	// 通过reflect获取num中的值
	pointer := reflect.ValueOf(&num)
	newValue := pointer.Elem()

	fmt.Println("type of pointer: ", newValue.Type())
	fmt.Println("settability of pointer: ", newValue.CanSet())

	newValue.SetFloat(77)
	fmt.Println("new value of pointer: ", newValue)
}
