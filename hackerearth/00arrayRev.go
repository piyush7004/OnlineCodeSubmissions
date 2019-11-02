// Write your code here
package main

import (
	"fmt"
	"strconv"
)

func main() {
	size := 0
	fmt.Scanf("%d", &size)
	arr := []int{}
	var inpT string
	for i := 0; i < size; i++ {
		fmt.Scanln(&inpT)
		// text := scanner.Text()
		inp, err := strconv.Atoi(inpT)
		if err != nil {
			fmt.Println("please provide proper number ", err)
		}
		arr = append(arr, inp)
	}
	for i := size - 1; i >= 0; i-- {
		fmt.Println(arr[i])
	}
}
