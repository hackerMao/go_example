package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func copyfile(src, dest string) {
	srcFile, err := os.Open(src)
	if err != nil {
		log.Fatalf("Failed to open file %s, error: %s\n", src, err)
	}
	defer srcFile.Close()

	descFile, err := os.Create(dest)
	if err != nil {
		log.Fatalf("Failed to open file %s, error: %s\n", dest, err)
	}
	defer descFile.Close()

	buf := make([]byte, 1024*1024)
	reader := bufio.NewReader(srcFile)
	writer := bufio.NewWriter(descFile)
	for {
		// n, err := srcFile.Read(buf)
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to read file %s, error : %s\n", src, err)
		}
		// descFile.Write(buf[:n])
		writer.Write(buf[:n])
		writer.Flush()
	}
}

func main() {
	src := flag.String("s", "", "src file")
	dest := flag.String("d", "", "desc file")
	help := flag.Bool("h", false, "help")

	flag.Usage = func() {
		fmt.Println(
			`
			Usage: copyfile -s srcfile -d descfile
			OPtions:
		`)
		flag.PrintDefaults()
	}

	flag.Parse()
	if *help || *src == "" || *dest == "" {
		flag.Usage()
	} else {
		copyfile(*src, *dest)
	}
}
