package main
import
	"fmt"

	func main() {
var anu Person
	anu.firstName = "Anurag"
	anu.lastName = "Raj"
	anu.age = 20
	fmt.Println("details", anu)

	
	}

	type Person struct {
		firstName string
		lastName  string
		age       int
	}