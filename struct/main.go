package main

import (
	"fmt"
)

/**
对于结构体来说，值接受者方法在golang中会自动创建指针接收者方法
所以结构体的方法方法既可以被值对象又可以被指针对象使用
*/

type Animal struct {
	Name     string
	Language string
}

func (a *Animal) Say() {
	fmt.Printf("%s说: %s\n", a.Name, a.Language)
}

func (a Animal) Play() {
	fmt.Printf("%s玩皮球\n", a.Name)
}

func main() {
	dog := Animal{"大黄", "汪汪汪"}
	fmt.Printf("%T\n", dog)
	dog.Say()
	dog.Play()

	duck := &Animal{"丑小鸭", "嘎嘎嘎"}
	fmt.Printf("%T\n", duck)
	duck.Say()
	duck.Play()
}
