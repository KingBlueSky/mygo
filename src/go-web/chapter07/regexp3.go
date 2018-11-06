package main

import (
	"fmt"
	"regexp"
)

func main() {
	a := "I am learning Go language"

	re, _ := regexp.Compile("[a-z]{2,4}")

	// 查找符合正则的第一个
	one := re.Find([]byte(a))
	fmt.Println("Find:", string(one))

	// 查找符合正则的所有slice，n小于0表示返回全部符合的字符串，不然就是返回指定的长度
	all := re.FindAll([]byte(a), -1)
	fmt.Println("FindAll:")
	for _, v := range all {
		fmt.Println(string(v))
	}

	// 查找符合条件的index位置，开始位置和结束位置
	index := re.FindIndex([]byte(a))
	fmt.Println("FindIndex", index)

	// 查找符合条件的所有的index位置，n同上
	allIndex := re.FindAllIndex([]byte(a), -1)
	fmt.Println("FindAllIndex", allIndex)

	re2, _ := regexp.Compile("am(.*)lang(.*)")

	// 查找Submatch,返回数组，第一个元素是匹配的全部元素，第二个元素是第一个()里面的，第三个是第二个()里面的
	// 下面的输出第一个元素是"am learning Go language"
	// 第二个元素是 "learning Go",注意包含空格的输出
	// 第三个元素是"uage"
	subMatch := re2.FindSubmatch([]byte(a))
	fmt.Println("FindSubMatch:")
	for _, v := range subMatch {
		fmt.Println(string(v))
	}
}
