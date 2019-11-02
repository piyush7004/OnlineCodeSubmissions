package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, _ := os.Open("input")
	defer file.Close()

	reader := bufio.NewReaderSize(file, 1024*1024)
	val, _, _ := reader.ReadLine()
	val, _, _ = reader.ReadLine()
	sarr := strings.Fields(string(val))
	arr := []int{}
	m := 0
	for _, s := range sarr {
		m, _ = strconv.Atoi(s)
		arr = append(arr, m)
	}
	stack := []int{}
	for _, v := range arr {
		stack = processStack(stack, v)
	}
	for _, v := range stack {
		fmt.Printf("%d ", v)
	}
}

func processStack(stack []int, item int) []int {
	if len(stack) == 0 {
		stack = append(stack, item)
		return stack
	}
	for item > stack[len(stack)-1] {
		stack = stack[:len(stack)-1]
		if len(stack) == 0 {
			stack = append(stack, item)
			return stack
		}
	}
	stack = append(stack, item)
	return stack
}
