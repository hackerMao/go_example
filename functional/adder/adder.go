package main

import "fmt"

func adder() func(a int) int {
	sum := 0
	return func(a int) int {
		sum += a
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	fmt.Println(adder()(2))
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + ... + %d = %d\n", i, a(i))
	}

	// 正统函数式编程
	a2 := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a2 = a2(i)
		fmt.Printf("0 + ... + %d = %d\n", i, s)
	}
}
