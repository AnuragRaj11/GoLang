package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// fmt.Println("Hello, what's your name?")
	var name string;
	// fmt.Scan(&name)
	// fmt.Println("Hello, " + name + "! How are you today?"
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ = reader.ReadString('\n')
	fmt.Println("Hello, " + name + "! How are you today?")
}