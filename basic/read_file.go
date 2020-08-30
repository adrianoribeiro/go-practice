package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	filename := os.Args[1]

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	text := string(contents)
	count := len(strings.Fields(text))

	fmt.Printf("There are %d words in this file.\n", count)
}
