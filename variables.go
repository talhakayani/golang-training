package main
import ("fmt")

func main() {
	var firstName  string = "Talha"
	var surName = "Kayani"
	var age int = 26
	height := 5.11

	fmt.Println("First Name: ", firstName)
	fmt.Println("Sur Name: ", surName)
	fmt.Println("Age: ", age)
	fmt.Println("Height: ", height)

	// Variable decleration without assign a value
	fmt.Println("\n\nVariable decleration without assign a value")
	var a int
	var b string
	var c bool

	fmt.Println("a: ", a) // 0
	fmt.Println("b: ", b) // empty string ""
	fmt.Println("c: ", c) // false

	// Value assignment after decleration 
	fmt.Println("\n\nValue assignment after decleration ")

	var v1 string
	var v2 int
	var v3 bool

	v1 = "Hello World!"
	v2 = 100
	v3 = true
	
	fmt.Println("v1: ", v1)
	fmt.Println("v2: ", v2)
	fmt.Println("v3: ", v3)


	// Multiple Variable declaration
	fmt.Println("\n\nMultiple Variable declaration")

	var a1, b1, c1, d1 int = 100, 10, 24, 30

	fmt.Println("a1, b1, c1, d1", a1, b1, c1, d1)


	// Different type of variables in single line

	var a2, b2, c2, d2 = "Talha", 26, 5.11, true
	a3, b3, c3 := "Kayani", 3.02, 280000

	fmt.Println("a2, b2, c2, d2", a2, b2, c2, d2)
	fmt.Println("a3, b3, c3", a3, b3, c3)
}