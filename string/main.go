package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// 判断字符串中的汉字的数量
	s1 := "hello 世界！"
	sum := hanCounts(s1)
	fmt.Printf("汉字的数量为：%d\n", sum)

	// how do you do 单词出现的次数
	s2 := "how do you do"
	m := charCounts(s2)
	for k, v := range m {
		fmt.Printf("%v 出现的次数：%d\n", k, v)
	}

	// 回文判断
	s3 := "黄山落叶松叶落山黄"
	if isPalindrome(s3) {
		fmt.Printf("%v 是回文\n", s3)
	} else {
		fmt.Printf("%v 不是回文\n", s3)
	}
}

func hanCounts(s string) int {
	sum := 0
	for _, v := range s {
		if unicode.Is(unicode.Han, v) {
			sum += 1
		}
	}
	return sum
}

func charCounts(s string) map[string]int {
	slice := strings.Split(s, " ")
	m := make(map[string]int, 10)
	for _, v := range slice {
		if count, ok := m[v]; ok {
			m[v] = count + 1
		} else {
			m[v] = 1
		}
	}
	return m
}

func isPalindrome(s string) bool {
	length := len(s)
	if length == 0 || length%2 == 0 {
		return false
	}
	r := make([]rune, 0, length)
	for _, v := range s {
		r = append(r, v)
	}
	fmt.Println(r)
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			return false
		}
	}
	return true
}
