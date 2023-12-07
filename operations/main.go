package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("Hello, Operations")

	//The use of the arithmetic operators
	{
		price, tax := 275.00, 27.40

		sum := price + tax 
		difference := price - tax
		product := price * tax
		quotient := price / tax

		fmt.Println(sum)
		fmt.Println(difference)
		fmt.Println(product)
		fmt.Println(quotient)
	}

	//An example of overflow of integer and floating-point numbers
	{
		var intVal = math.MaxInt64
		var floatVal = math.MaxFloat64

		fmt.Println(intVal * 2)
		fmt.Println(floatVal * 2)
		fmt.Println(math.IsInf((floatVal * 2), 0))
	}

	//Go provides the % operator, which returns the remainder when one integer value is divided by another
	//It can return negative values
	{
		posResult := 3 % 2
		negResult := -3 % 2
		absResult := math.Abs(float64(negResult))

		fmt.Println(posResult)
		fmt.Println(negResult)
		fmt.Println(absResult)
	}

	//Go provides a set of operators for incrementing and decrementing numeric values
	//These operators can be applied to integer and floating-point numbers
	{
		value := 10.2
		value++
		fmt.Println(value)
		value += 2
		fmt.Println(value)
		value -= 2
		fmt.Println(value)
		value--
		fmt.Println(value)
	}

	//The + operator can be used to concatenate strings to produce longer strings
	{
		greeting := "Hello"
		language := "Go"
		combinedString := greeting + ", " + language

		fmt.Println(combinedString)
	}

	//The values used with the comparison operators must all be of the same type, 
	//or they must be untyped constants that can be represented as the target type
	{
		first := 100
		const second = 200.00

		equal := first == second
		notEqual := first != second
		lessThan := first < second
		lessThanOrEqual := first <= second
		greaterThan := first > second
		greaterThanOrEqual := first >= second

		fmt.Println(equal)
		fmt.Println(notEqual)
		fmt.Println(lessThan)
		fmt.Println(lessThanOrEqual)
		fmt.Println(greaterThan)
		fmt.Println(greaterThanOrEqual)
	}

	//Pointers can be compared to see if they point at the same memory location
	{
		first := 100

		second := &first
		third := &first

		alpha := 100
		beta := &alpha

		fmt.Println(second == third)
		fmt.Println(second == beta)
	}

	//If you want to compare values, then you should follow the pointers
	{
		first := 100

		second := &first
		third := &first

		alpha := 100
		beta := &alpha

		fmt.Println(*second == *third)
		fmt.Println(*second == *beta)
	}

	//Examples of the logical operators being used to produce values that are assigned to variables
	{
		maxMph := 50
		passengerCapacity := 4
		airbags := true

		familyCar := passengerCapacity > 2 && airbags
		sportsCar := maxMph > 100 || passengerCapacity == 2
		canCategorize := !familyCar && !sportsCar

		fmt.Println(familyCar)
		fmt.Println(sportsCar)
		fmt.Println(canCategorize)
	}
	//Go doesn’t allow types to be mixed in operations and 
	//will not automatically convert types, except in the case of untyped constants
	{
		// kayak := 275
		// soccerBall := 19.50

		// total := kayak + soccerBall

		// fmt.Println(total)
		
		// # command-line-arguments
		// .\main.go:140:12: invalid operation: kayak + soccerBall (mismatched types int and float64)
	}

	//An explicit conversion transforms a value to change its type
	{
		kayak := 275
		soccerBall := 19.50

		total := float64(kayak) + soccerBall

		fmt.Println(total)
	}

	//Explicit conversions can be used only when the value can be represented in the target type
	//You can convert between numeric types and between strings and runes
	//Combinations, such as converting int values to bool values, are not supported
	//Care must be taken when choosing the values to convert
	//Explicit conversions can cause a loss of precision in numeric values or cause overflows
	{
		kayak := 275
		soccerBall := 19.50
		total := kayak + int(soccerBall)
		fmt.Println(total)
		fmt.Println(int8(total))
	}

	//These functions return float64 values
	{
		fmt.Println(math.Ceil(27.1))
		fmt.Println(math.Floor(27.1))
		fmt.Println(math.Round(27.1))
		fmt.Println(math.RoundToEven(27.1))

		//Which can then be explicitly converted to the int type
		kayak := 275
		soccerBall := 19.50

		total := kayak + int(math.Round(soccerBall))

		fmt.Println(total)
	}

	//Parsing from strings
	{
		val1 := "true"
		val2 := "false"
		val3 := "not true"

		bool1, b1err := strconv.ParseBool(val1)
		bool2, b2err := strconv.ParseBool(val2)
		bool3, b3err := strconv.ParseBool(val3)

		fmt.Println("Bool 1", bool1, b1err)
		fmt.Println("Bool 2", bool2, b2err)
		fmt.Println("Bool 3", bool3, b3err)
	}

	//Checking for an Error 
	{
		val1 := "0"
		
		bool1, b1err := strconv.ParseBool(val1)

		if b1err == nil {
			fmt.Println("Parsed value:", bool1)
		} else {
			fmt.Println("Cannot parse", val1)
		}
	}

	//Checking an Error in a Single Statement
	{
		val1 := "0"

		if bool1, b1err := strconv.ParseBool(val1); b1err == nil {
			fmt.Println("Parsed value:", bool1)
		} else {
			fmt.Println("Cannot parse", val1)
 		}
	}

	//Parsing an Integer
	{
		val1 := "100"
		//the first argument is the value, the second - is the base and the third is the size
		int1, int1err := strconv.ParseInt(val1, 0, 8)

		if int1err == nil {
			fmt.Println("Parsed value:", int1)
		} else {
			fmt.Println("Cannot parse", val1)
		}
	}
	//The size only specifies the data size that the parsed value must be
    //able to fit into
	{
		val1 := "500"

		int1, int1err := strconv.ParseInt(val1, 0, 8)

		if int1err == nil {
			fmt.Println("Parsed value:", int1)
		} else {
			fmt.Println("Cannot parse", val1, int1err)
		}
	}

	//Specifying a size of 8 when calling the ParseInt function allows me to perform an explicit conversion
	//to the int8 type without the possibility of overflow
	{
		val1 := "100"

		int1, int1err := strconv.ParseInt(val1, 0, 8)

		if int1err == nil {
			smallInt := int8(int1)
			fmt.Println("Parsed value:", smallInt)
		} else {
			fmt.Println("Cannot parse", val1, int1err)
		}
	}

	//Parsing a Binary Value 
	{
		val1 := "100"

		//a base of 2 means the string will be interpreted as a binary value
		int1, int1err := strconv.ParseInt(val1, 2, 8)

		if int1err == nil {
			smallInt := int8(int1)
			fmt.Println("Parsed value:", smallInt)
		} else {
			fmt.Println("Cannot parse", val1, int1err)
		}
	}

	//You can leave the Parse<Type> functions to detect the base for a value using a prefix
	{
		val1 := "0b1100100"

 		int1, int1err := strconv.ParseInt(val1, 0, 8)
		
		if int1err == nil {
			smallInt := int8(int1)
			fmt.Println("Parsed value:", smallInt)
		} else {
			fmt.Println("Cannot parse", val1, int1err)
		}
	}

	//For many projects, the most common parsing task is to create int values 
	//from strings that contain decimal numbers
	{
		val1 := "100"

		int1, int1err := strconv.ParseInt(val1, 10, 0)

		if int1err == nil {
			var intResult int = int(int1)
			fmt.Println("Parsed value:", intResult)
		} else {
			fmt.Println("Cannot parse", val1, int1err)
		}
	}

	//This is such a common task that the strconv package provides the Atoi function, 
	//which handles the parsing and explicit conversion in a single step
	{
		val1 := "100"

		int1, int1err := strconv.Atoi(val1)

		if int1err == nil {
			var intResult int = int1
			fmt.Println("Parsed value:", intResult)
		} else {
			fmt.Println("Cannot parse", val1, int1err)
		}
	}

	//The ParseFloat function is used to parse strings containing floating-point numbers
	{
		val1 := "48.95"

		float1, float1err := strconv.ParseFloat(val1, 64)

		if float1err == nil {
			fmt.Println("Parsed value:", float1)
		} else {
			fmt.Println("Cannot parse:", val1, float1err)
		}
	}

	//The ParseFloat function can parse values expressed with an exponent
	{
		val1 := "4.895e+01"

		float1, float1err := strconv.ParseFloat(val1, 64)

		if float1err == nil {
			fmt.Println("Parsed value:", float1)
		} else {
			fmt.Println("Cannot parse", val1, float1err)
		}
	}

	//The FormatBool function accepts a bool value and returns a string representation
	{
		val1 := true
		val2 := false

		str1 := strconv.FormatBool(val1)
		str2 := strconv.FormatBool(val2)

		fmt.Println("Formatted value 1: " + str1)
		fmt.Println("Formatted value 2: " + str2)
	}

	//The FormatInt and FormatUint functions format integer values as strings
	{
		val := 275

		base10String := strconv.FormatInt(int64(val), 10)
		base2String := strconv.FormatInt(int64(val), 2)

		fmt.Println("Base 10: " + base10String)
		fmt.Println("Base 2: " + base2String)
	}

	//Integer values are most commonly represented using the int type 
	//and are converted to strings using base 10
	//The strconv package provides the Itoa function, which is 
	//a more convenient way to perform this specific conversion
	{
		val := 275

		base10String := strconv.Itoa(val)
		base2String := strconv.FormatInt(int64(val), 2)

		fmt.Println("Base 10: " + base10String)
		fmt.Println("Base 2: " + base2String)
	}

	//Expressing floating-point values as strings requires additional configuration options 
	//because different formats are available
	{
		val := 49.95

		Fstring := strconv.FormatFloat(val, 'f', 2, 64)
		Estring := strconv.FormatFloat(val, 'e', -1, 64)

		fmt.Println("Format F: " + Fstring)
		fmt.Println("Format E: " + Estring)
	}
}