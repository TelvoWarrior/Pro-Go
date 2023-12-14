package main

import "fmt"

func main()  {
	fmt.Println("Hello, Functions")
	
	fmt.Println("About to call function")
	printPrice()
	fmt.Println("Function complete")

	//Parameters allow a function to receive data values when it is called, allowing its behavior to be altered.
	fmt.Println("\nParameters allow a function to receive data values when it is called, allowing its behavior to be altered:")
	printPriceTwo("Kayak", 275, 0.2)
	printPriceTwo("Lifejacket", 48.95, 0.2)
	printPriceTwo("Soccer Ball", 19.50, 0.15)

	//The type can be omitted when adjacent parameters have the same type
	fmt.Println("\nThe type can be omitted when adjacent parameters have the same type:")
	printPriceThree("Kayak", 275, 0.2)
	printPriceThree("Lifejacket", 48.95, 0.2)
	printPriceThree("Soccer Ball", 19.50, 0.15)

	//Omitting Parameter Names
	fmt.Println("\nOmitting Parameter Names:")
	printPriceFour("Kayak", 275, 0.2)
	printPriceFour("Lifejacket", 48.95, 0.2)
	printPriceFour("Soccer Ball", 19.50, 0.15)
	
	//Functions can also omit names from all their parameters
	fmt.Println("\nFunctions can also omit names from all their parameters:")
	printPriceFive("Kayak", 275, 0.2)
	printPriceFive("Lifejacket", 48.95, 0.2)
	printPriceFive("Soccer Ball", 19.50, 0.15)

	//To understand the issue that variadic parameters solve, it can be helpful to consider the alternative
	fmt.Println("\nTo understand the issue that variadic parameters solve, it can be helpful to consider the alternative:")
	printSuppliers("Kayak", []string {"Acme Kayaks", "Bob's Boats", "Crazy Canoes"})
 	printSuppliers("Lifejacket", []string {"Sail Safe Co"})

	//Variadic parameters allow a function to receive a variable number of arguments more elegantly
	fmt.Println("\nVariadic parameters allow a function to receive a variable number of arguments more elegantly:")
	printSuppliersTwo("Kayak", "Acme Kayaks", "Bob's Boats", "Crazy Canoes")
 	printSuppliersTwo("Lifejacket", "Sail Safe Co")
}

//Functions are groups of statements that can be used and reused as a single action
func printPrice()  {
	kayakPrice := 275.00
	kayakTax := kayakPrice * 0.2
	fmt.Println("Price:", kayakPrice, "Tax:", kayakTax)
}

//Parameters allow a function to receive data values when it is called, allowing its behavior to be altered.
func printPriceTwo(product string, price float64, taxRate float64) {
	taxAmount := price * taxRate
	fmt.Println(product, "price:", price, "Tax:", taxAmount)
}

//The type can be omitted when adjacent parameters have the same type
func printPriceThree(product string, price, taxRate float64) {
	taxAmount := price * taxRate
	fmt.Println(product, "price:", price, "Tax:", taxAmount)
}

//Omitting Parameter Names
func printPriceFour(product string, price, _ float64) {
	taxAmount := price * 0.25
	fmt.Println(product, "price:", price, "Tax:", taxAmount)
}

//Functions can also omit names from all their parameters
func printPriceFive(string, float64, float64) {
	// taxAmount := price * 0.25
	fmt.Println("No parameters")
}

//To understand the issue that variadic parameters solve, it can be helpful to consider the alternative
func printSuppliers(product string, suppliers []string) {
	for _, supplier := range suppliers {
		fmt.Println("Product:", product, "Supplier:", supplier)
	}
}

//Variadic parameters allow a function to receive a variable number of arguments more elegantly
func printSuppliersTwo(product string, suppliers ...string) {
	for _, supplier := range suppliers {
		fmt.Println("Product:", product, "Supplier:", supplier)
	}
	//The variadic parameter must be the last parameter defined by the function, 
	//and only a single type can be used, such as the string type in this example
}