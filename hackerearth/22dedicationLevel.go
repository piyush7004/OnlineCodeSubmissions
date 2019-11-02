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
	bytes, _, _ := reader.ReadLine()
	strarr := []string{}
	var max, m, i, v int
	var s, k string
	var err error
	nmap := make(map[string]int)
	for {
		bytes, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		strarr = strings.Fields(string(bytes))
		m, _ = strconv.Atoi(strarr[1])
		nmap[strarr[0]] = m
	}
	for i < 3 {
		for k, v = range nmap {
			if v > max {
				max = v
				s = k
			}
		}
		fmt.Println(s)
		max = 0
		delete(nmap, s)
		i++
	}
}
