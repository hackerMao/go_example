package main

import (
	"fmt"
)

func main() {
	arr := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range arr {
		sum += v
	}
	fmt.Printf("sum(arr) : %d\n", sum)

	// 找出和為8的兩個下標
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i]+arr[j] == 8 {
				fmt.Printf("(%d, %d)\n", i, j)
			}
		}
	}
}
