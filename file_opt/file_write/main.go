package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func writeDemo1() {
	file, err := os.OpenFile("./test.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:  %+v\n", err)
	}
	defer file.Close()
	file.Write([]byte("醉舞经阁半卷书\n"))
	file.WriteString("坐井说天阔\n")
	file.WriteString("大志戏功名\n海斗量福祸")

}

func writeDemo2() {
	file, err := os.OpenFile("./test.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:  %+v\n", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		_, err := writer.WriteString(fmt.Sprintf("number %d\n", i))
		if err != nil {
			return
		}
	}
	writer.Flush()
}

func writeDemo3() {
	str := "hello kunming"
	err := ioutil.WriteFile("./test.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file filed, error：", err)
		return
	}
}

func main() {
	//writeDemo1()
	//writeDemo2()
	writeDemo3()
}
