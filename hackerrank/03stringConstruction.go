package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()
	reader := bufio.NewReaderSize(file, 1024*1024)
	bytes, _, _ := reader.ReadLine()
	var err error
	for {
		bytes, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(processConstruction(string(bytes)))
	}
}

func processConstruction(s string) int {
	nmap := make(map[string]int)
	var cost int
	var found bool
	for _, c := range s {
		_, found = nmap[string(c)]
		if found {
			cost++
		} else {
			nmap[string(c)] = 1
		}
		fmt.Println(nmap, found)
	}
	return cost
}
