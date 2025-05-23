package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 10
	var data float64 = float64(num)
	data = data+1.5
	str:=strconv.Itoa(num)
	fmt.Println("num:", num)
	fmt.Println("data:", data)
	fmt.Println("str:", str)
	str="100"
	numb,err:=strconv.Atoi(str)
	if err!=nil{
		fmt.Println("Error:", err)
	}
	numb=numb+1
	fmt.Println("numb:", numb)


}