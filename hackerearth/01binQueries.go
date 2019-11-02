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
	size := 0
	file, _ := os.Open("input")
	defer file.Close()

	strArr := []string{}
	reader := bufio.NewReaderSize(file, 10*1024*1024)
	var line []byte
	var err error
	for {
		line, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		}
		strArr = append(strArr, string(line))
	}
	size, _ = strconv.Atoi(strArr[0])
	arr := []int{}
	sarr := strings.Fields(strArr[1])

	for i := 0; i < size; i++ {
		inp, err := strconv.Atoi(sarr[i])
		if err != nil {
			fmt.Println("Please provide the proper integer as input")
		}
		arr = append(arr, inp)
	}
	nmap := make(map[int]int)
	for _, v := range arr {
		nmap[v]++
	}
	qarr := []int{}
	queries, _ := strconv.Atoi(strArr[2])

	for i := 3; i < len(strArr); i++ {
		// fmt.Scan(&inp)
		inp, err := strconv.Atoi(strArr[i])
		if err != nil {
			fmt.Println("Please provide the proper integer as input")
		}
		qarr = append(qarr, inp)
	}
	for i := 0; i < queries; i++ {
		val, ok := nmap[qarr[i]]
		if ok {
			fmt.Println(val)
		} else {
			fmt.Println("NOT PRESENT")
		}

	}
}
