package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	err  error
	num  int
	line string
	pair []string
)

func main() {
	for {
		if line, err = scanLine(); err != nil {
			break
		} else {
			pair = strings.Split(line, ",")
			fmt.Println(pair[0] + "," + pair[1])
		}
	}
	if err != nil && err != io.EOF {
		eHandle("cityMapper: something wrong when scan the input")
	}
	return
}

func eHandle(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func scanLine() (string, error) {
	var words []rune
	var word rune
	var err error = nil
	for {
		_, err = fmt.Scanf("%c", &word)
		if err != nil || word == '\n' {
			break
		}
		words = append(words, word)
	}
	if err == io.EOF && len(words) > 0 {
		err = nil
	}
	return string(words), err
}
