package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) ReflectCallFunc() {
	fmt.Println("Allen.Wu ReflectCallFunc")
}

func main() {
	user := User{1, "Allen.Wu", 27}
	DoFiledAndMethod(user)
}

func DoFiledAndMethod(input interface{}) {

	// 获取变量的类型
	getType := reflect.TypeOf(input)
	fmt.Println("get type is: ", getType.Name())

	// 获取变量的值
	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is: ", getValue)

	// 遍历所有的field
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 遍历所有 method
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}
