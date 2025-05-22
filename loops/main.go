package main
import
	"fmt"

	func main() {
		x:=4
		if x%2==0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}

		day:=4
		switch day {
		case 1:
			fmt.Println("Monday")
			case 2:
			fmt.Println("Tuesday")
			case 3:
			fmt.Println("Wednesday")
			case 4:
			fmt.Println("Thursday")
			case 5:
			fmt.Println("Friday")
		}

		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}