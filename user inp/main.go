package main

import "fmt"

func main() {

	fmt.Println("Hello, what's your name?")
	var name string;
	fmt.Scan(&name)
	fmt.Println("Hello, " + name + "! How are you today?")
}