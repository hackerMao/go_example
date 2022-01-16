package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func write() {
	file, err := os.Create("user.csv")
	if err != nil {
		log.Fatalf("Failed to open file user.csv, error: %s\n", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Write([]string{"编号", "名字", "性别"})
	writer.Write([]string{"1", "hah", "女"})
	writer.Write([]string{"2", "heh", "男"})
	writer.Flush()
}

func read() {
	f, err := os.Open("user.csv")
	if err != nil {
		log.Fatalf("Failed to open file user.csv, error: %s\n", err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to read file user.csv, error: %s\n", err)
		}
		fmt.Println(record)
	}
}

func main() {
	// write()
	read()
}
