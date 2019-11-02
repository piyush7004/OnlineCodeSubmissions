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
	i := 0
	defer file.Close()
	reader := bufio.NewReaderSize(file, 1024*1024)
	val, _, _ := reader.ReadLine()
	strarr := strings.Fields(string(val))
	k, _ := strconv.Atoi(strarr[1])
	val, _, _ = reader.ReadLine()
	strarr = strings.Fields(string(val))
	k = k % len(strarr)
	s := ""
	out := []string{}
	for i = 0; i < len(strarr); i++ {
		s = strarr[(i+k)%len(strarr)]
		out = append(out, s)
	}
	fmt.Println(strings.Join(out, " "))
}
