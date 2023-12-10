package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Hello, Collections")

	//Go arrays are a fixed length and contain elements of a single type, which are accessed by index
	{
		var names [3]string

		names[0] = "Kayak"
		names[1] = "Lifejacket"
		names[2] = "Paddle"

		fmt.Println(names)
	}

	//Arrays can be defined and populated in a single statement using the literal syntax
	{
		names := [3]string{"Kayak", "Lifejacket", "Paddle"}
		fmt.Println(names)
	}

	//The type of an array is the combination of its size and underlying type
	{
		// names := [3]string { "Kayak", "Lifejacket", "Paddle" }

		// var otherArray [4]string = names

		// fmt.Println(otherArray)
		//.\main.go:29:30: cannot use names (variable of type [3]string) as [4]string value in variable declaration
	}

	//Assigning an array to a new variable copies the array and copies the values it contains
	{
		names := [3]string{"Kayak", "Lifejacket", "Paddle"}

		otherArray := names

		names[0] = "Canoe"

		fmt.Println("names:", names)
		fmt.Println("otherArray:", otherArray)
	}

	//A pointer can be used to create a reference to an array
	{
		names := [3]string{"Kayak", "Lifejacket", "Paddle"}

		otherArray := &names

		names[0] = "Canoe"

		fmt.Println("names:", names)
		fmt.Println("otherArray:", *otherArray)
	}

	//The comparison operators == and != can be applied to arrays
	{
		names := [3]string{"Kayak", "Lifejacket", "Paddle"}
		moreNames := [3]string{"Kayak", "Lifejacket", "Paddle"}

		same := names == moreNames

		fmt.Println("comparison:", same)
		//Arrays are equal if they are of the same type and contain equal elements in the same order.
	}

	//Arrays are enumerated using the for and range keywords
	{
		names := [3]string{"Kayak", "Lifejacket", "Paddle"}

		for index, value := range names {
			fmt.Println("Index:", index, "Value:", value)
		}
	}

	//Go doesn’t allow variables to be defined and not used.
	//If you don’t need both the index and the value, you can use
	//an underscore (the _ character) instead of a variable name
	{
		names := [3]string{"Kayak", "Lifejacket", "Paddle"}

		for _, value := range names {
			fmt.Println("Value:", value)
		}
	}

	//The best way to think of slices is as a variable-length array
	//One way to define a slice is to use the built-in make function
	{
		names := make([]string, 3)

		names[0] = "Kayak"
		names[1] = "Lifejacket"
		names[2] = "Paddle"

		fmt.Println(names)
		//The make function accepts arguments that specify the type and length of the slice
	}

	//Slices can also be created using a literal syntax
	{
		names := []string{"Kayak", "Lifejacket", "Paddle"}

		fmt.Println(names)
	}

	//One of the key advantages of slices is that they can be expanded to accommodate additional elements
	{
		names := []string{"Kayak", "Lifejacket", "Paddle"}

		names = append(names, "Hat", "Gloves")

		fmt.Println(names)
		//The built-in append function accepts a slice and
		//one or more elements to add to the slice, separated by commas
	}

	//The original slice—and its backing array—still exists and can be used
	{
		names := []string{"Kayak", "Lifejacket", "Paddle"}

		appendedNames := append(names, "Hat", "Gloves")

		names[0] = "Canoe"

		fmt.Println("names:", names)
		fmt.Println("appendedNames:", appendedNames)
	}

	//Creating and copying arrays can be inefficient.
	//If you expect that you will need to append items to a slice,
	//you can specify additional capacity when using the make function
	{
		names := make([]string, 3, 6)

		names[0] = "Kayak"
		names[1] = "Lifejacket"
		names[2] = "Paddle"

		fmt.Println("len:", len(names))
		fmt.Println("cap:", cap(names))
	}

	//The underlying array isn’t replaced when the append function is called on
	//a slice with enough capacity to accommodate the new elements
	{
		names := make([]string, 3, 6)

		names[0] = "Kayak"
		names[1] = "Lifejacket"
		names[2] = "Paddle"

		appendedNames := append(names, "Hat", "Gloves")

		names[0] = "Canoe"

		fmt.Println("names:", names)
		fmt.Println("appendedNames:", appendedNames)
	}

	//The append function can be used to append one slice to another
	{
		names := make([]string, 3, 6)

		names[0] = "Kayak"
		names[1] = "Lifejacket"
		names[2] = "Paddle"

		moreNames := []string{"Hat", "Gloves"}

		appendedNames := append(names, moreNames...)

		fmt.Println("appendedNames:", appendedNames)
	}

	//Slices can be created using existing arrays
	{
		products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}

		someNames := products[1:3]
		allNames := products[:]

		fmt.Println("someNames:", someNames)
		fmt.Println("allNames:", allNames)
	}

	//Displaying Slice Length and Capacity
	{
		products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}

		someNames := products[1:3]
		allNames := products[:]

		fmt.Println("someNames:", someNames)
		fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
		fmt.Println("allNames", allNames)
		fmt.Println("allNames len", len(allNames), "cap:", cap(allNames))
	}

	//Appending an Element to a Slice
	{
		products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}

		someNames := products[1:3]
		allNames := products[:]

		someNames = append(someNames, "Gloves")

		fmt.Println("someNames:", someNames)
		fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
		fmt.Println("allNames", allNames)
		fmt.Println("allNames len", len(allNames), "cap:", cap(allNames))
	}

	//Appending Another Element
	{
		products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}

		someNames := products[1:3]
		allNames := products[:]

		someNames = append(someNames, "Gloves")
		someNames = append(someNames, "Boots")

		fmt.Println("someNames:", someNames)
		fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
		fmt.Println("allNames", allNames)
		fmt.Println("allNames len", len(allNames), "cap:", cap(allNames))
	}

	//Ranges can include a maximum capacity, which provides some degree of control
	//over when arrays will be duplicated
	{
		products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}

		someNames := products[1:3:3]
		allNames := products[:]

		someNames = append(someNames, "Gloves")

		fmt.Println("someNames:", someNames)
		fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
		fmt.Println("allNames", allNames)
		fmt.Println("allNames len", len(allNames), "cap:", cap(allNames))
	}

	//Slices can also be created from other slices, although the relationship
	//between slices isn’t preserved if they are resized
	{
		products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}

		allNames := products[1:]
		someNames := allNames[1:3]

		allNames = append(allNames, "Gloves")
		allNames[1] = "Canoe"

		fmt.Println("someNames:", someNames)
		fmt.Println("allNames:", allNames)
	}

	//The copy function can be used to duplicate an existing slice,
	//selecting some or all the elements but ensuring that
	//the new slice is backed by its own array
	{
		products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}

		allNames := products[1:]
		someNames := make([]string, 2)
		copy(someNames, allNames)
		//The copy function accepts two arguments, which are the destination slice and the source slice
	}

	//A common pitfall is to try to copy elements into a slice that has not been initialized
	{
		products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}

		allNames := products[1:]
		var someNames []string
		copy(someNames, allNames)

		fmt.Println("allNames:", allNames)
		fmt.Println("someNames:", someNames)
		//No elements have been copied to the destination slice. This happens because uninitialized slices have
		//zero length and zero capacity. The copy function stops copying when the length of the destination length
		//is reached, and since the length is zero, no copying occurs
	}

	//Fine-grained control over the elements that are copied can be achieved using ranges
	{
		products := [4]string { "Kayak", "Lifejacket", "Paddle", "Hat" }

		allNames := products[1:]
		someNames := []string { "Boots", "Canoe" }
		copy(someNames[1:], allNames[2:3])
		
		fmt.Println("someNames:", someNames)
		fmt.Println("allNames", allNames)
	}

	//If the destination slice is larger than the source slice, 
	//then copying will continue until the last element in the source has been copied
	{
		products := []string { "Kayak", "Lifejacket", "Paddle", "Hat" }
		replacementProducts := []string { "Canoe", "Boots" }
		
		copy(products, replacementProducts)
		
		fmt.Println("products:", products)
	}

	//If the destination slice is smaller than the source slice, 
	//then copying continues until all the elements in the destination slice have been replaced
	{
		products := []string { "Kayak", "Lifejacket", "Paddle", "Hat"}
		replacementProducts := []string { "Canoe", "Boots"}

		copy(products[0:1], replacementProducts)

		fmt.Println("products:", products)
		//The range used for the destination creates a slice with length of one, which means that only one
		//element will be copied from the source array
	}

	//There is no built-in function for deleting slice elements, 
	//but this operation can be performed using the ranges and the append function
	{
		products := [4]string { "Kayak", "Lifejacket", "Paddle", "Hat" }

		deleted := append(products[:2], products[3:]...)
		fmt.Println("Deleted:", deleted)
	}

	//Slices are enumerated in the same way as arrays, with the for and range keywords
	{
		products := []string { "Kayak", "Lifejacket", "Paddle", "Hat" }

		for index, value := range products[2:] {
			fmt.Println("Index:", index, "Value:", value)
		}
	}

	//There is no built-in support for sorting slices, but 
	//the standard library includes the sort package, which
	//defines functions for sorting different types of slice
	{
		products := []string { "Kayak", "Lifejacket", "Paddle", "Hat" }

		sort.Strings(products)

		for index, value := range products {
			fmt.Println("Index:", index, "Value:", value)
		}
	}

	//Go restricts the use of the comparison operator so that slices can be compared only to the nil value.
	{
		// p1 := []string { "Kayak", "Lifejacket", "Paddle", "Hat" }

		// p2 := p1
		// fmt.Println("Equal:", p1 == p2)
		//.\main.go:367:25: invalid operation: p1 == p2 (slice can only be compared to nil)
	}

	//There is one way to compare slices, however. 
	//The standard library includes a package named reflect, 
	//which includes a convenience function named DeepEqual.
	{
		p1 := []string { "Kayak", "Lifejacket", "Paddle", "Hat" }

		p2 := p1

		fmt.Println("Equal:", reflect.DeepEqual(p1, p2))
	}

	//Getting the Array Underlying a Slice
	{
		p1 := []string { "Kayak", "Lifejacket", "Paddle", "Hat"}
		arrayPtr := (*[3]string)(p1)
		array := *arrayPtr

		fmt.Println(array)
	}

	//Maps are a built-in data structure that associates data values with keys.
	//Unlike arrays, where values are associated with sequential integer locations, 
	//maps can use other data types as keys
	{
		products := make(map[string]float64, 10)

		products["Kayak"] = 279
		products["Lifejacket"] = 48.95

		fmt.Println("Map size:", len(products))
		fmt.Println("Price:", products["Kayak"])
		fmt.Println("Price:", products["Hat"])
	}

	//Maps can also be defined using a literal syntax
	{
		products := map[string]float64 {
			"Kayak": 279,
			"Lifejacket": 48.95,
		}

		fmt.Println("Map size:", len(products))
		fmt.Println("Price:", products["Kayak"])
		fmt.Println("Price:", products["Hat"])
	}

	//Maps return the zero value for the value type when reads are performed 
	//for which there is no key. This can make it difficult to differentiate 
	//between a stored value that happens to be the zero value and a nonexistent key
	{
		products := map[string]float64 {

			"Kayak" : 279,
			"Lifejacket": 48.95,
			"Hat": 0,
		}

		fmt.Println("Hat:", products["Hat"])
	}

	//Maps produce two values when reading a value
	{
		products := map[string]float64 {

			"Kayak" : 279,
			"Lifejacket": 48.95,
			"Hat": 0,
		}

		value, ok := products["Hat"]

		if ok {
			fmt.Println("Stored value:", value)
		} else {
			fmt.Println("No stored value")
		}
		//This is known as the “comma ok” technique, where values are assigned to 
		//two variables when reading a value from a map
	}

	//This technique can be streamlined using an initialization statement
	{
		products := map[string]float64 {

			"Kayak" : 279,
			"Lifejacket": 48.95,
			"Hat": 0,
		}

		if value, ok := products["Hat"]; ok {
			fmt.Println("Stored value:", value)
		} else {
			fmt.Println("No stored value")
		}
	}

	//Items are removed from the map using the built-in delete function
	{
		products := map[string]float64 {

			"Kayak" : 279,
			"Lifejacket": 48.95,
			"Hat": 0,
		}

		delete(products, "Hat")

		if value, ok := products["Hat"]; ok {
			fmt.Println("Stored value:", value)
		} else {
			fmt.Println("No stored value")
		}
	}

	//Maps are enumerated using the for and range keywords
	{
		products := map[string]float64 {

			"Kayak" : 279,
			"Lifejacket": 48.95,
			"Hat": 0,
		}

		for key, value := range products {
			fmt.Println("Key:", key, "Value:", value)
		}
	}

	//Enumerating a Map in Order
	{
		products := map[string]float64 {

			"Kayak" : 279,
			"Lifejacket": 48.95,
			"Hat": 0,
		}

		keys := make([]string, 0, len(products))
		for key, _ := range products {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			fmt.Println("Key:", key, "Value:", products[key])
		}
	}

	//Go treats strings as arrays of bytes and supports the array index and slice range notation
	{
		var price string = "$48.95"

		var currency byte = price[0]
		var amountString string = price[1:]
		amount, parseErr := strconv.ParseFloat(amountString, 64)

		fmt.Println("Currency:", currency)
		if parseErr == nil {
			fmt.Println("Amount:", amount)
		} else {
			fmt.Println("Parse Error:", parseErr)
		}
	}

	//Slicing a string produces another string, but an explicit conversion 
	//is required to interpret the byte as the character it represents
	{
		var price string = "$48.95"

		var currency string = string(price[0])
		var amountString string = price[1:]
		amount, parseErr := strconv.ParseFloat(amountString, 64)

		fmt.Println("Currency:", currency)
		if parseErr == nil {
			fmt.Println("Amount:", amount)
		} else {
			fmt.Println("Parse Error:", parseErr)
		}
	}

	//It it contains a pitfall, which can be seen if the currency symbol is changed
	{
		var price string = "€48.95"

		var currency string = string(price[0])
		var amountString string = price[1:]
		amount, parseErr := strconv.ParseFloat(amountString, 64)

		fmt.Println("Currency:", currency)
		if parseErr == nil {
			fmt.Println("Amount:", amount)
		} else {
			fmt.Println("Parse Error:", parseErr)
			//Parse Error: strconv.ParseFloat: parsing "\x82\xac48.95": invalid syntax
		}
	}

	//The rune type represents a Unicode code point, which is essentially a single character. 
	//To avoid slicing strings in the middle of characters, 
	//an explicit conversion to a rune slice can be performed
	{
		var price []rune = []rune("€48.95")
		var currency string = string(price[0])
		var amountString string = string(price[1:])
		amount, parseErr := strconv.ParseFloat(amountString, 64)

		fmt.Println("Currency:", currency)
		if parseErr == nil {
			fmt.Println("Amount:", amount)
		} else {
			fmt.Println("Parse Error:", parseErr)
		}
	}

	//A for loop can be used to enumerate the contents of a string. 
	//This feature shows some clever aspects of the way that Go deals with the mapping of bytes to runes.
	{
		var price = "€48.95"

		for index, char := range price {
			fmt.Println(index, char, string(char))
		}
	}

	//If you want to enumerate the underlying bytes without them being converted to characters, 
	//then you can perform an explicit conversion to a byte slice
	{
		var price = "€48.95"

		for index, char := range []byte(price) {
			fmt.Println(index, char)
		}
	}
}
