package main

import 
	"fmt"

func main() {
	fmt.Println("Hello Go!")
	test()
	ans:= math(1, 4)
	fmt.Println(ans)
}

func test(){
	fmt.Println("new test")
}
func math(a,b int) (res int ) {
	return a + b
}