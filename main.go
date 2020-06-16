package main

import (
	"fmt"
	"os"

	"github.com/sebastianwebber/test4dummies/style"
)

var sentence string

func init() {
	if len(os.Args) > 1 {
		sentence = os.Args[1]
	}
}

func main() {
	pretty, err := style.Sentence(sentence)

	if err != nil {
		fmt.Printf("fail: %v\n", err)
		return
	}
	fmt.Println(pretty)
}
