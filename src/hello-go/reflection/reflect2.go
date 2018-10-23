package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num = 1.2345
	fmt.Println("type: ", reflect.TypeOf(num))
	fmt.Println("value: ", reflect.ValueOf(num))

	// 已知变量的类型
	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)
	fmt.Println(pointer.Interface().(*float64))
	fmt.Println(value.Interface().(float64))
}
