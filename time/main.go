package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%T\n", now)
	fmt.Printf("%v\n", now)

	// 时间格式化
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:05"))

	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano()) //纳秒

	// 字符串转时间
	t, _ := time.Parse("2006/01/02 15:04:05", "2022/01/13 19:08:23")
	fmt.Println(t)
	fmt.Println(t.Year())
	fmt.Println(t.Month())

	// 时间计算
	time.Sleep(time.Second * 3)
	endTime := time.Now()
	sub := endTime.Sub(now)
	fmt.Printf("%T, %v\n", sub, sub)
}
