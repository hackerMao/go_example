package main

import "fmt"

type Sender interface {
	Send(to string, msg string) error
	SendAll(to []string, msg string) error
}

type EmailSender struct {
}

func (e *EmailSender) Send(to string, msg string) error {
	fmt.Printf("发送邮件到：%s, 邮件内容：%s\n", to, msg)
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
