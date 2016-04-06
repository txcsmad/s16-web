// Every .go file is part of a package
// Runnable packages have to be in a package named main.
package main

import (
	"fmt" // built-in formatting package
)

// Executable Go programs have a main function
func main() {
	// • Names starting with uppercase letters (for example: Println) are "public" or exported.
	// • Names starting with lowercase letters are "private" or unexported and cannot
	//   be accessed outside the package that defines them.
	// • Within the same package, all names (uppercase or otherwise) are accessible.
	// • As you can see from the use of the "©" in the string, Go has excellent Unicode support.
	fmt.Println("hello, world ©")

	// Create a car c1 and drive it 100 miles:
	c1 := Car{
		wheels:   8,
		distance: 0,
		color:    "white",
	}
	c1.drive(100)
	fmt.Println(c1.distance) // 100

	// The above is exactly the same as doing:
	var c2 Car = Car{
		8, 0, "white",
	}
	c2.drive(100)
	fmt.Println(c2.distance) // 100
}

// • The func keyword declares a function.
// • The variable types come after the variable name.
// • If the parameters have the same type, you can save some keystrokes and mention
//   the type just once (as was the case in the parameters here).
// • The return type comes after the parameters list.
func add(a, b int) int {
	return a + b
}

// • Car is a custom type we defined. It has 3 fields: wheels, distance, color.
// • The capitalization rule for exported v. unexported fields also applies to struct fields as well.
// • In this case, all 3 fields are unexported since they start with a lowercase character.
type Car struct {
	wheels, distance int
	color            string
}

// • drive is a method attached to the Car type.
// • We can attach a method to a type by mentioning the type (in this case, *Car)
//   in front of the function name.
// • The Go convetion is to use short names for arguments.
func (c *Car) drive(d int) {
	c.distance += d
}
