package main

import "fmt"

func main() {
	num,_ := getData()
	fmt.Println("Number:", num)

}

func getData() (int, string) {
	return 11, "Anurag"
}
