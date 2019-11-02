package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	reader := bufio.NewReaderSize(file, 1024*1024)
	bytes, _, _ := reader.ReadLine()
	s := string(bytes)
	fmt.Println(constructNow(s))
}

func constructNow(s string) int {
	nmap := make(map[byte]int)
	for _, b := range []byte(s) {
		nmap[b]++
	}
	return len(nmap)
}
