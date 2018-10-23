package main

import "fmt"

func main() {
	num := []int{2, 4, 7, 6, 10, 2}
	target := 12
	result := TwoSum1(num, target)
	fmt.Println(result)
	result2 := TwoSum2(num, target)
	fmt.Println(result2)
}

func TwoSum1(num []int, target int) []int {
	n := len(num)

	for i, v := range num {
		for j := i + 1; j < n; j++ {
			if v+num[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func TwoSum2(num []int, target int) []int {
	m := make(map[int]int, len(num))

	for i, v := range num {
		sub := target - v
		if j, ok := m[sub]; ok {
			return []int{j, i}
		} else {
			m[v] = i
		}
	}

	return nil
}
