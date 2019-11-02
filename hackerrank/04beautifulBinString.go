package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input")
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	reader := bufio.NewReaderSize(file, 1024*1024)
	bytes, _, _ := reader.ReadLine()
	bytes, _, _ = reader.ReadLine()
	st := string(bytes)
	fmt.Println((len(st) - len(strings.ReplaceAll(st, "010", ""))) / 3)
	fmt.Println((len(st) - len(strings.Replace(st, "010", "", -1))) / 3)
}

// func myReplaceAll(s, source, target string) string {

// 	for _, b := range []byte(s) {

// 	}

// }
