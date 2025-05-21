package main

import
	"fmt"

	func main() {

		num:= [5]int{1, 2, 3, 4}
		fmt.Println("Slice:", num)
		fmt.Println("Length of slice:", len(num))
		fmt.Println("Capacity of slice:", cap(num))

		nums:=make([]int, 3, 5)
		nums=append(nums, 1)
		fmt.Println("Slice:", nums)
		fmt.Println("Length of slice:", len(nums))
		fmt.Println("Capacity of slice:", cap(nums))
	}