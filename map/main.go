package main
import
	"fmt"

	func main() {
		students := make(map[string]int)
		students["Alice"] = 90
		students["Bob"] = 85
		students["Charlie"] = 95
		students["David"] = 80

		fmt.Println("Students and their scores:", students["Alice"])
		fmt.Println("Students and their scores:", students["Bob"])
		fmt.Println("Students and their scores:", students["Charlie"])

		delete(students, "David")
		fmt.Println("Students and their scores:", students["David"])
		
	}