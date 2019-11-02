package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()
	reader := bufio.NewReaderSize(file, 1024*1024)
	bytes, _, _ := reader.ReadLine()
	bytes, _, _ = reader.ReadLine()
	strarr := strings.Fields(string(bytes))
	nmap := make(map[byte]int)
	var v int

	for _, s := range strarr {
		bytes = removeDuplicates([]byte(s))
		for _, b := range bytes {
			nmap[b]++
		}
	}
	max := 0
	for _, v = range nmap {
		if v > max {
			max = v
		}
	}
	fmt.Println(max)
}
func removeDuplicates(bytes []byte) []byte {
	nmap := make(map[byte]int)
	var found bool
	out := []byte{}
	for _, b := range bytes {
		_, found = nmap[b]
		if !found {
			nmap[b]++
			out = append(out, b)
		}
	}
	return out
}
