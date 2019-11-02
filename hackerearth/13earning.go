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
	s := ""
	strarr := []string{}
	var m, rate, count int
	stack := []int{}
	var err error
	for {
		val, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		count++
		strarr = strings.Fields(string(val))
		if count == 1 {
			rate, _ = strconv.Atoi(strarr[1])
		}
		if count == 2 {
			//get array
			for _, s = range strarr {
				m, _ = strconv.Atoi(s)
				stack = processStack(stack, m)
			}
			fmt.Println(rate * len(stack))

			count = 0
			stack = nil
		}
	}
}

func processStack(stack []int, v int) []int {
	if len(stack) == 0 {
		stack = append(stack, v)
		return stack
	}
	if v <= stack[len(stack)-1] {
		return stack
	} else {
		stack = append(stack, v)
		return stack
	}
}
