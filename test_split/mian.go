package main

import (
	"fmt"
	"github.com/go_example/split_string"
)

func main() {
	split := split_string.Split("babcef", "b")
	fmt.Printf("%#v\n", split)
}
