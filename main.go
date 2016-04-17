package main

import (
	"fmt"
	"strconv"
)

func inc1(i int) {
	i++
	fmt.Println("inc1: i =  " + strconv.Itoa(i))
}

func inc2(i *int) {
	*i++
	fmt.Println("inc2: i =  " + strconv.Itoa(*i))
}

func main() {
	num := 10
	inc1(num)                                // inc1: i = 11
	fmt.Println("i =  " + strconv.Itoa(num)) // num = 10

	inc2(&num)
	fmt.Println("i =　" + strconv.Itoa(num)) // num = 11
}
