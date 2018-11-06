package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// 从标准输入流读取数据
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Please input your name:")
	// 读取数据直到碰到\n为止
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
		os.Exit(1)
	} else {
		name := input[:len(input)-1]
		fmt.Printf("Hello, %s! What can I do for you\n", name)
	}

	for {
		input, err = inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
			continue
		}
		input = input[:len(input)-1]
		// 转换成小写
		input = strings.ToLower(input)

		switch input {
		case "":
			continue
		case "nothing", "byte":
			fmt.Println("Bye!")
			os.Exit(0)
		default:
			fmt.Println("Sorry, I didn't catch you.")
		}
	}
}
