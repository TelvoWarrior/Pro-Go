package main

import (
	"fmt"
	"sort"
	// "math/rand"
)

func main() {
	//Literal values
	fmt.Println("Hello, Go")
	fmt.Println(20+20)
	fmt.Println(20+30)

	//Typed constants
	//Typed constants are defined using the const keyword, followed by a name, a type, and a value assignment
	{
		const price float32 = 275.00
		const tax float32 = 27.50
		fmt.Println(price+tax)
	}

	//The untyped constant feature makes constants easier to work with 
	//because the Go compiler will perform limited automatic conversion
	{
		const price float32 = 275.00
		const tax float32 = 27.50
		const quantity = 2
		fmt.Println("Total:", quantity * (price + tax))
	}

	//A single statement can be used to define several constants
	{
		const price, tax float32 = 275, 27.50
		const quantity, inStock = 2, true
		fmt.Println("Total:", quantity * (price+tax))
		fmt.Println("In stock: ", inStock)
	}

	//Literal values are untyped constants
	{
		const price, tax float32 = 275, 27.50
		const quantity, inStock = 2, true
		fmt.Println("Total:", 2 * quantity * (price + tax))
		fmt.Println("In stock: ", inStock)
	}

	//Variables are defined using the var keyword, and, unlike constants, 
	//the value assigned to a variable can be changed
	{
		var price float32 = 275.00
		var tax float32 = 27.50
		fmt.Println(price + tax)
		price = 300
		fmt.Println(price + tax)
	}

	//The Go compiler can infer the type of variables based on the initial value, 
	//which allows the type to be omitted
	{
		var price = 275.00
		var price2 = price
		fmt.Println(price)
		fmt.Println(price2)
	}

	//Omitting a type doesn’t have the same effect for variables as it does for constants, 
	//and the Go compiler will not allow different types to be mixed
	{
		// var price = 275.00
		// var tax float32 = 27.50
		// fmt.Println(price + tax)
		// invalid operation: price + tax (mismatched types float64 and float32)
	}

	//Variables can be defined without an initial value
	{
		var price float32
		fmt.Println(price)
		price = 275.00
		fmt.Println(price)
	}

	//A single statement can be used to define several variables
	{
		var price, tax = 275.00, 27.50
		fmt.Println(price + tax)
	}

	//A type must be specified if no initial values are assigned,
	//and all variables will be created using the specified type and assigned their zero value.
	{
		var price, tax float64 
		price = 275.00
		tax = 27.50
		fmt.Println(price + tax)
	}

	//The short variable declaration provides a shorthand for declaring variables
	{
		price := 275.00
		fmt.Println(price)
	}

	//Multiple variables can be defined with a single statement 
	//by creating comma-separated lists of names and values
	{
		price, tax, inStock := 275.00, 27.50, true
		fmt.Println("Total:", price + tax)
		fmt.Println("In stock:", inStock)
	}

	//Redefining a variable is allowed if the short syntax is used, as long as at least one 
	//of the other variables being defined doesn’t already exist and the type of the variable
	//doesn’t change.
	{
		price, tax, inStock := 275.00, 27.50, true
		fmt.Println("Total:", price + tax)
		fmt.Println("In stock:", inStock)
		price2, tax := 200.00, 25.00
		fmt.Println("Total 2:", price2 + tax)
	}

	//It is illegal in Go to define a variable and not use it
	{
		// price, tax, inStock, discount := 275.00, 27.50, true, true
		// var salesPerson = "Alice"
		// fmt.Println("Total:", price + tax)
		// fmt.Println("In stock:", inStock)

		// .\main.go:125:24: discount declared but not used
		// .\main.go:126:7: salesPerson declared but not used
	}

	//For these situations, Go provides the blank identifier, 
	//which is used to denote a value that won’t be used
	{
		price, tax, inStock, _ := 275.00, 27.50, true, true
		var _ = "Alice"
		fmt.Println("Total:", price + tax)
		fmt.Println("In stock:", inStock)
	}

	//To understand how pointers work, the best place to start is 
	//understanding what Go does when pointers are not used
	{
		first := 100
		second := first

		first++

		fmt.Println("First:", first)
		fmt.Println("Second:", second)
	}

	//A pointer is a variable whose value is a memory address
	{
		first := 100
		var second *int = &first

		first++

		fmt.Println("First:", first)
		fmt.Println("Second:", second)
	}

	//The phrase following a pointer means reading the value at the memory address that the pointer refers
	//to, and it is done using an asterisk (the * character)
	{
		first := 100
		second := &first

		first++

		fmt.Println("First:", first)
		fmt.Println("Second:", *second)
	}

	//Following a Pointer and Changing the Value 
	{
		first := 100
		second := &first

		first++
		*second++

		fmt.Println("First:", first)
		fmt.Println("Second:", *second)
	}

	//The second use of a pointer, which is to use it as a value in its own right and
	//assign it to another variable
	{
		first := 100
		second := &first

		first++
		*second++

		var myNewPointer *int
		myNewPointer = second
		*myNewPointer++

		fmt.Println("First:", first)
 		fmt.Println("Second:", *second)
	}

	//Pointers that are defined but not assigned a value have the zero-value nil
	{
		first := 100
		var second *int

		fmt.Println(second)
		second = &first
		fmt.Println(second)
	}

	//A runtime error will occur if you follow a pointer that has not been assigned a value
	{
		// first := 100
		// var second *int
		// fmt.Println(*second)
		// second = &first
		// fmt.Println(second == nil)

		// panic: runtime error: invalid memory address or nil pointer dereference
		// [signal 0xc0000005 code=0x0 addr=0x0 pc=0x29f42d]

		// goroutine 1 [running]:
		// main.main()
		// 		C:/Users/aliak/Desktop/Pro-Go/basicFeatures/main.go:221 +0xf2d
		// exit status 2
	}

	//It is possible to create a pointer whose value is the memory
	//address of another pointer
	{
		first := 100
		second := &first
		third := &second

		fmt.Println(first)
		fmt.Println(*second)
		fmt.Println(**third)
	}

	//An example of when working with values is useful
	{
		//Create an array of three string values
		names := [3]string{"Alice", "Charlie", "Bob"}

		//Assign the values in position 1 to a variable called secondName
		secondName := names[1]

		//Write to the console the value of the secondName variable
		fmt.Println(secondName)

		//Sort the array
		sort.Strings(names[:])

		//Write to the console the value of the secondName variable again
		fmt.Println(secondName)
	}

	//An example of using a pointer instead of a value
	{
		names := [3]string{"Alice", "Charlie", "Bob"}
		secondPosition := &names[1]
		fmt.Println(*secondPosition)
		sort.Strings(names[:])
		fmt.Println(*secondPosition)
	}
}