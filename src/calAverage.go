package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	Mapper []string
	curKey string
	Val    float64
	curNum int64
	curAmt float64
	line   string
	err    error
)

func main() {
	for {
		if line, err = scanLine(); err != nil {
			break
		}
		if Mapper = strings.Split(line, ","); len(Mapper) < 2 {
			fmt.Println(Mapper, len(Mapper))
			eHandle("calAverage: input format error")
		}
		Val, err := strconv.ParseFloat(Mapper[1], 64)
		if err != nil {
			break
		}

		if Mapper[0] == curKey {
			curAmt += Val
			curNum++
		} else {
			if curAmt != 0 {
				fmt.Println(curKey + "\t" + strconv.FormatFloat(curAmt/float64(curNum), 'f', -1, 64))
			}
			curKey = Mapper[0]
			curNum = 1
			curAmt = Val
		}
	}

	if err != io.EOF {
		eHandle("calAverage: something wrong")
	} else {
		fmt.Println(curKey + "\t" + strconv.FormatFloat(curAmt/float64(curNum), 'f', -1, 64))
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
