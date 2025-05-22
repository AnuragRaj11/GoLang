package main

import "fmt"

func main() {
	var num int
	num = 2

	// var ptr *int
	// ptr = &num

	ptr:=&num
	fmt.Println("value of num", num)
	fmt.Println("address of num",ptr)
	fmt.Println("value of num using pointer", *ptr)
}
