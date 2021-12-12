package main

import "fmt"

/**
--一切源于变量是值拷贝和go语言中函数的return不是原子操作，在底层分两步执行--
	第一步： 返回值赋值
	执行defer：函数中存在defer，则defer执行的时机是在第一步和第二步之间
	第二步：真正的RET指令
*/
func f1() int {
	x := 5
	defer func() {
		x++ // 修改的是X不是返回值
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 1、返回值赋值为x 5赋值给x, 2、x++, 3、RET指令
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // 1、返回值赋值为y = x = 5, 2、defer修改的是x, 3、RET指令
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 // 1、返回值赋值为x=5, 2、把x的副本5传进去, 3、RET指令
}

func f5() (x int) {
	defer func(x int) int {
		x++
		return x
	}(x)
	return 5 // 1、返回值赋值为x=5, 2、把x的副本5传进去 但匿名函数返回并没有被接收, 3、RET指令
}

func f6() (x int) {
	defer func(x *int) {
		*x++
	}(&x)
	return 5 // 1、返回值赋值为x=5, 2、把x的地址传进去 此时x的值被修改, 3、RET指令
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	//fmt.Println(f1())
	//fmt.Println(f2())
	//fmt.Println(f3())
	//fmt.Println(f4())
	//fmt.Println(f5())
	//fmt.Println(f6())
	//
	//a := 10
	//b := &a
	//*b++
	//fmt.Printf("a : %d\n", a)
	//fmt.Printf("b : %v\n", b)
	//fmt.Printf("&a : %v\n", &a)

	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1

	// 1. a := 1
	// 2. b := 2
	// 3. defer calc("1", 1, calc("10", 1, 2))
	// 4. calc("10", 1, 2) // "10" 1 2 3
	// 5, defer calc("1", 1, 3) // "1" 1 3 4
	// 6. a = 0
	// 7. defer calc("2", 0, calc("20", 0, 2))
	// 8. calc("20", 0, 2) // "20" 0 2 2
	// 9. defer calc("2", 0, 2) // "2" 0 2 2
	// 10 b = 1
}
