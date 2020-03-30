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
			fmt.Println(pair[0] + "\t" + pair[1])
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
	var bits []byte
	var ch byte
	var err error = nil
	for {
		_, err = fmt.Scanf("%c", &ch)
		if err != nil || ch == '\n' {
			break
		}
		bits = append(bits, ch)
	}
	return string(bits), err
}

