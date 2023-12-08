package main

import (
	"fmt"
	"strconv"
)

func main() {
	kayakPrice := 275.00
	fmt.Println("Price:", kayakPrice)

	//An if statement is used to execute a group of statements only 
	//when a specified expression produces the bool value true when it is evaluated
	{
		kayakPrice := 275.00

		if kayakPrice > 100 {
			fmt.Println("Price is greater than 100")
		}
	}

	//The else keyword can be used to create additional clauses in an if statement
	{
		kayakPrice := 275.00

		if kayakPrice > 500 {
			fmt.Println("Price is greater than 500")
		} else if kayakPrice < 300 {
			fmt.Println("Price is less than 300")
		}
	}

	//Defining Multiple else/if Clauses
	{
		kayakPrice := 275.00

		if kayakPrice > 500 {
			fmt.Println("Price is greater than 500")
		} else if kayakPrice < 100 {
			fmt.Println("Price is less than 100")
		} else if kayakPrice > 200 && kayakPrice < 300 {
			fmt.Println("Price is between 200 and 300")
		}
	}

	//The else keyword can also be used to create a fallback clause, 
	//whose statements will be executed only if all 
	//the if and else/if expressions in the statement produce false results
	{
		kayakPrice := 275.00

		if (kayakPrice > 500) {
			fmt.Println("Price is greater than 500")
		} else if (kayakPrice < 100) {
			fmt.Println("Price is less than 100")
		} else {
			fmt.Println("Price not matched by earlier expressions")
		}
	}

	//Each clause in an if statement has its own scope
	{
		kayakPrice := 275.00

		if kayakPrice > 500 {
			scopedVar := 500
			fmt.Println("Price is greater than", scopedVar)
		} else if kayakPrice < 100 {
			scopedVar := "Price is less than 100"
			fmt.Println(scopedVar)
		} else {
			scopedVar := false
			fmt.Println("Matched: ", scopedVar)
		}
	}

	//Go allows an if statement to use an initialization statement, 
	//which is executed before the if statement’s expression is evaluated
	{
		priceString := "275"

		if kayakPrice, err := strconv.Atoi(priceString); err == nil {
			fmt.Println("Price:", kayakPrice)
		} else {
			fmt.Println("Error:", err)
		}
	}

	//The for keyword is used to create loops that repeatedly execute statements
	{
		counter := 0
		for {
			fmt.Println("Counter:", counter)
			counter++
			if counter > 3 {
				break
			}
		}
	}

	//A common requirement is to repeat until a condition is reached. 
	//This is such a common requirement that the condition 
	//can be incorporated into the loop syntax
	{
		counter := 0
		for counter <= 3 {
			fmt.Println("Counter:", counter)
			counter++
		}
	}

	//Loops can be defined with additional statements that are executed 
	//before the first iteration of the loop (known as the initialization statement) 
	//and after each iteration (the post statement)
	{
		for counter := 0; counter <= 3; counter++ {
			fmt.Println("Counter:", counter)
		}
	}

	//The continue keyword can be used to terminate the execution of the for loop’s statements 
	//for the current value and move to the next iteration
	{
		for counter := 0; counter <= 3; counter++ {
			if counter == 1 {
				continue
			}
			fmt.Println("Counter:", counter)
		}
	}

	//The for keyword can be used with the range keyword to create loops that enumerate over sequences
	{
		product := "Kayak"

		for index, character := range product {
			fmt.Println("Index:", index, "Character:", string(character))
		}
	}
	
	//Go will report an error if a variable is defined but not used
	//You can omit the value variable from the for... range statement 
	//if you require only the index values
	{
		product := "Kayak"

		for index := range product {
			fmt.Println("Index:", index)
		}
	}

	//The blank identifier can be used when you require only the values 
	//in the sequence and not the indices
	{
		product := "Kayak"

		for _, character := range product {
			fmt.Println("Character:", string(character))
		}
	}

	//The range keyword can also be used with the built-in data structures 
	//that Go provides — arrays, slices, and maps
	{
		products := []string {"Kayak", "Lifejacket", "Soccer Ball"}

		for index, element := range products {
			fmt.Println("Index:", index, "Element:", element)
		}
	}

	//A switch statement provides an alternative way to control execution flow
	{
		product := "Kayak"

		for index, character := range product {
			switch (character) {
			case 'K':
				fmt.Println("K at position:", index)
			case 'y':
				fmt.Println("y at position:", index)
			}
		}
	}

	//Go switch statements do not fall through automatically, but 
	//multiple values can be specified with a comma-separated list
	{
		product := "Kayak"

		for index, character := range product {
			switch character {
			case 'K', 'k':
				fmt.Println("K or k at position", index)
			case 'y':
				fmt.Println("y at position", index)
			}
		}
	}

	//Although the break keyword isn’t required to terminate every case statement, 
	//it can be used to end the execution of statements before 
	//the end of the case statement is reache
	{
		product := "Kayak"

		for index, character := range product {
			switch character {
			case 'K', 'k':
				if character == 'k' {
					fmt.Println("Lowercase k at position", index)
					break
				}
				fmt.Println("Uppercase K at position", index)
			case 'y':
				fmt.Println("y at position", index)
			}
		}
	}

	//Go switch statements don’t automatically fall through, but 
	//this behavior can be enabled using the fallthrough keyword
	{
		product := "Kayak"

		for index, character := range product {
			switch character {
			case 'K':
				fmt.Println("Uppercase character")
				fallthrough
			case 'k':
				fmt.Println("k at position", index)
			case 'y':
				fmt.Println("y at position", index)
			}
		}
	}

	//The default keyword is used to define a clause that will be 
	//executed when none of the case statements matches the switch statement’s value
	{
		product := "Kayak"

		for index, character := range product {
			switch (character) {
			case 'K', 'k':
				if (character == 'k') {
					fmt.Println("Lowercase k at position", index)
					break
				}
				fmt.Println("Uppercase K at position", index)
			case 'y':
				fmt.Println("y at position", index)
			default:
				fmt.Println("Character", string(character), "at position", index)
			}
		}
	}

	//a common problem in switch statements where an expression 
	//is used to produce the comparison value
	{
		for counter := 0; counter < 20; counter++ {
			switch(counter / 2) {
			case 2, 3, 5, 7:
				fmt.Println("Prime value:", counter / 2)
			default:
				fmt.Println("Non-prime value:", counter / 2)
			}
		}
	}

	//The duplication can be avoided using an initialization statement
	{
		for counter := 0; counter < 20; counter++ {
			switch val := counter / 2; val {
			case 2, 3, 5, 7:
				fmt.Println("Prime value:", val)
			default:
				fmt.Println("Non-prime value:", val)
			}
		}
	}

	//Omitting the comparison value and using expressions in the case statements
	{
		for counter := 0; counter < 10; counter++ {
			switch {
			case counter == 0:
					fmt.Println("Zero value")
			case counter < 3:
				fmt.Println(counter, "is < 3")
			case counter >= 3 && counter < 7:
				fmt.Println(counter, "is >= 3 && < 7")
			default:
				fmt.Println(counter, "is >= 7")
			}
		}
	}

	//Label statements allow execution to jump to a different point, 
	//giving greater flexibility than other flow control features
	{
		counter := 0
		target: fmt.Println("Counter", counter)
		counter++
		if counter < 5 {
			goto target
		}
	}
}