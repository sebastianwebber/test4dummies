package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	inputName string
)

func parseEnvs() {
	if name := os.Getenv("NAME"); name != "" {
		inputName = name
	}
}

func init() {
	flag.StringVar(&inputName, "name", "World", "Put your name here!")
	flag.Parse()
	parseEnvs()
}

func main() {
	fmt.Printf("Hello %s!\n", inputName)
}

// superUpper its the best and only great string upper function! =p
func superUpper(nome string) (out string, err error) {

	if nome == "" {
		err = fmt.Errorf("Empty string not allowed here")
		return
	}

	out = strings.ToUpper(nome)

	return
}
