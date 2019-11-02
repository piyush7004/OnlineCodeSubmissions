package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	reader := bufio.NewReaderSize(file, 1024*1024)
	bytes, _, _ := reader.ReadLine()
	strarr := strings.Fields(string(bytes))
	arr := []int32{}
	for _, s := range strarr {
		m, _ := strconv.Atoi(s)
		arr = append(arr, int32(m))
	}
	fmt.Println(pickNum(arr))
}

func pickNum(arr []int32) int32 {
	nmap := make(map[int32]int)
	for _, v := range arr {
		nmap[v]++
	}
	max := 0
	var a, b int
	for k, _ := range nmap {
		a = nmap[k]
		b = nmap[k-1]
		if (a + b) > max {
			max = a + b
		}
	}
	return int32(max)
}
