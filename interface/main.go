package main

import "fmt"

/**
结构体方法接收指针，则可将结构体指针对象赋值给接口，接口可以直接调用方法
结构体方法接收值类型，既可以将指针也可以将值类型赋值给接口，因为在使用值类型时go会自动使用引用，这在go里算是一个语法糖
值接收者会自动生成指针接收者，所以值对象和指针对象都可以赋值给接口，而指针接收者则不能
*/

type Sender interface {
	Send(to string, msg string) error
	SendAll(to []string, msg string) error
}

type EmailSender struct {
}

func (e *EmailSender) Send(to string, msg string) error {
	fmt.Printf("发送邮件到: %s, 邮件内容: %s\n", to, msg)
	return nil
}

func (e *EmailSender) SendAll(to []string, msg string) error {
	for _, v := range to {
		e.Send(v, msg)
	}
	return nil
}

func main() {
	var sender Sender = &EmailSender{}
	fmt.Printf("%T, %v\n", sender, sender)

	sender.Send("hacker_murray@163.com", "早上好")
	sender.SendAll([]string{"hacker_murray@163.com", "1765785706@qq.com"}, "早上好")
}
