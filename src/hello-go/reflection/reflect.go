package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x = 3.4
	fmt.Println("type: ", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("type: ", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value: ", v.Float())

	p := reflect.ValueOf(&x)
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())
	s := p.Elem()
	fmt.Println("settability of s:", s.CanSet())

	s.SetFloat(7.1)
	fmt.Println(s.Interface())
	fmt.Println(x)
}
