package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()
	reader := bufio.NewReaderSize(file, 1024*1024)
	val, _, _ := reader.ReadLine()
	var err error
	var count, i, k int
	s := ""
	strarr := []string{}
	ostrarr := []string{}

	for {
		val, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		count++
		strarr = strings.Fields(string(val))
		if count == 1 {
			k, _ = strconv.Atoi(strarr[1])
		}
		if count == 2 {
			k = k % len(strarr)
			//form ostrarr by rotation
			for i = 0; i < len(strarr); i++ {
				fmt.Println(k, len(strarr))
				s = strarr[(i+len(strarr)-k)%len(strarr)]
				ostrarr = append(ostrarr, s)
			}
			//print ostrarr
			fmt.Println(strings.Join(ostrarr, " "))
			count = 0
			ostrarr = nil

		}
	}

}
