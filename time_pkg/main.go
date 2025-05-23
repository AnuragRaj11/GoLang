package main

import (
	"fmt"
	"time"
)

func main() {
	currt:=time.Now()
	fmt.Println("Current Time:", currt)

	format:=currt.Format("2006-01-02 15:04:05")
	fmt.Println("Formatted Time:", format)

	layout:= "2006-01-02 15:04:05"
	str:="2023-10-01 12:00:00"
	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed Time:", t)
	}



}