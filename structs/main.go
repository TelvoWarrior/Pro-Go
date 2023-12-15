package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main()  {
	fmt.Println("Hello, Structs!")

	type Product struct{
		name, category string
		price float64
	}

	kayak := Product {
		name: "Kayak",
		category: "Watersports",
		price: 275,
	}

	fmt.Println(kayak.name, kayak.category, kayak.price)
	kayak.price = 300
	fmt.Println("Changed price:", kayak.price)

	//Values do not have to be provided for all fields when creating a struct value
	{
		type Product struct{
			name, category string
			price float64
		}
	
		kayak := Product {
			name: "Kayak",
			category: "Watersports",
		}
	
		fmt.Println(kayak.name, kayak.category, kayak.price)
		kayak.price = 300
		fmt.Println("Changed price:", kayak.price)
	}

	//The zero types are assigned to all fields if you define a struct-typed variable but don’t assign a value to it
	{
		type Product struct {
			name, category string
			price float64
		}

		var lifejacket Product

		fmt.Println("Name is zero value:", lifejacket.name == "")
		fmt.Println("Category is zero value:", lifejacket.category == "")
		fmt.Println("Price is zero value:", lifejacket.price == 0)
	}

	//Struct values can be defined without using names, 
	//as long as the types of the values correspond to the order
	//in which fields are defined by the struct type
	{
		type Product struct {
			name, category string
			price float64
		}

		var kayak = Product { "Kayak", "Watersports", 275.00 }
		fmt.Println("Name:", kayak.name)
		fmt.Println("Category:", kayak.category)
		fmt.Println("Price:", kayak.price)
	}

	//If a field is defined without a name, it is known as an embedded field, 
	//and it is accessed using the name of its type
	{
		type Product struct {
			name, category string
			price float64
		}

		type StockLevel struct {
			Product
			count int
		}

		stockItem := StockLevel {
			Product: Product{ "Kayak", "Watersports", 275 },
			count: 100,
		}

		fmt.Println("Name:", stockItem.Product.name)
		fmt.Println("Count:", stockItem.count)
	}

	//As noted earlier, field names must be unique with the struct type,
	//which means that you can define only one embedded field for a specific type.
	// If you need to define two fields of the same type, 
	//then you will need to assign a name to one of them
	{
		type Product struct {
			name, category string
			price float64
		}

		type StockLevel struct {
			Product
			Alternate Product
			count int
		}

		stockItem := StockLevel {
			Product: Product{ "Kayak", "Watersports", 275 },
			Alternate: Product{ "Lifejacket", "Watersports", 48.95 },
			count: 100,
		}

		fmt.Println("Name:", stockItem.Product.name)
		fmt.Println("Alt Name:", stockItem.Alternate.name)
	}

	//Struct values are comparable if all their fields can be compared
	{
		type Product struct {
			name, category string
			price float64
		}
	
		p1 := Product { name: "Kayak", category: "Watersports", price: 275.00 }
		p2 := Product { name: "Kayak", category: "Watersports", price: 275.00 }
		p3 := Product { name: "Kayak", category: "Boats", price: 275.00 }
	
		fmt.Println("p1 == p2:", p1 == p2)
		fmt.Println("p1 == p3:", p1 == p3)
	}

	//Structs cannot be compared if the struct type defines fields with incomparable types, such as slices
	{
		// type Product struct {
		// 	name, category string
		// 	price float64
		// 	otherNames []string
		// }
	
		// p1 := Product { name: "Kayak", category: "Watersports", price: 275.00 }
		// p2 := Product { name: "Kayak", category: "Watersports", price: 275.00 }
		// p3 := Product { name: "Kayak", category: "Boats", price: 275.00 }
	
		// fmt.Println("p1 == p2:", p1 == p2)
		// fmt.Println("p1 == p3:", p1 == p3)
		//.\main.go:144:28: invalid operation: p1 == p2 (struct containing []string cannot be compared)
		//.\main.go:145:28: invalid operation: p1 == p3 (struct containing []string cannot be compared)
	}

	//A struct type can be converted into any other struct type that has the same fields, 
	//meaning all the fields have the same name and type and are defined in the same order
	{
		type Item struct {
			name string
			category string
			price float64
		}

		prod := Product { name: "Kayak", category: "Watersports", price: 275.00 }
		item := Item { name: "Kayak", category: "Watersports", price: 275.00 }
		
		fmt.Println("prod == item:", prod == Product(item))
	}

	//Anonymous struct types are defined without using a name
	{
		writeName := func (val struct {
			name, category string
			price float64}) {
				fmt.Println("Name:", val.name)
			}
		
		type Product struct {
			name, category string
			price float64
		}

		type Item struct {
			name string
			category string
			price float64
		}

		prod := Product{ name: "Kayak", category: "Watersports", price: 275.00 }
		item := Item { name: "Stadium", category: "Soccer", price: 75000 }

		writeName(prod)
		writeName(item)
	}

	//Another example anonymous struct use
	{
		type Product struct {
			name, category string
			price float64
		}

		prod := Product { "Kayak", "Watersports", 275 }

		var builder strings.Builder
		json.NewEncoder(&builder).Encode(struct {
			ProductName string
			ProductCategory string
			ProductPrice float64
		}{
			ProductName: prod.name,
			ProductCategory: prod.category,
			ProductPrice: prod.price,
		})
		fmt.Println(builder.String())
	}

	//The struct type can be omitted when populating arrays, slices, and maps with struct values
	{
		type Product struct {
			name, category string
			price float64
		}

		type StockLevel struct {
			Product
			Alternate Product
			count int
		}

		array := [1]StockLevel {
			{
				Product: Product{ "Kayak", "Watersports", 275.00 },
				Alternate: Product{ "Lifejacket", "Watersports", 48.95 },
				count: 100,
			},
		}
		fmt.Println("Array:", array[0].Product.name)

		slice := []StockLevel {
			{
				Product: Product{ "Kayak", "Watersports", 275.00 },
				Alternate: Product{ "Lifejacket", "Watersports", 48.95 },
				count: 100,
			},
		}
		fmt.Println("Slice:", slice[0].Product.name)

		kvp := map[string]StockLevel {
			"kayak": {
				Product: Product{ "Kayak", "Watersports", 275.00 },
				Alternate: Product{ "Lifejacket", "Watersports", 48.95 },
				count: 100,
			},
		}
		fmt.Println("Map:", kvp["kayak"].Product.name)
	}

	//Assigning a struct to a new variable or using a struct as a function parameter 
	//creates a new value that copies the field values
	{
		type Product struct {
			name, category string
			price float64
		}

		p1 := Product {
			name: "Kayak",
			category: "Watersports",
			price: 275,
		}

		p2 := p1

		p1.name = "Original Kayak"

		fmt.Println("p1:", p1.name)
		fmt.Println("p2:", p2.name)
	}

	//References to struct values can be created using pointers
	{
		type Product struct {
			name, category string
			price float64
		}

		p1 := Product {
			name: "Kayak",
			category: "Watersports",
			price: 275,
		}

		p2 := &p1

		p1.name = "Original Kayak"

		fmt.Println("p1:", p1.name)
		fmt.Println("p2:", (*p2).name)
	}

	//Understanding the Struct Pointer Convenience Syntax
	{
		type Product struct {
			name, category string
			price float64
		}

		calcTax := func (product *Product) {
			if (*product).price > 100 {
				(*product).price += (*product).price * 0.2 
			}
		}

		kayak := Product { "Kayak", "Watersports", 275 }

		calcTax(&kayak)

		fmt.Println("Name:", kayak.name, "Category:", kayak.category, "Price:", kayak.price)
	}

	//Example of follow pointers to struct fields without needing an asterisk character
	{
		type Product struct {
			name, category string
			price float64
		}

		calcTax := func (product *Product) {
			if product.price > 100 {
				product.price += product.price * 0.2 
			}
		}

		kayak := Product { "Kayak", "Watersports", 275 }

		calcTax(&kayak)

		fmt.Println("Name:", kayak.name, "Category:", kayak.category, "Price:", kayak.price)
	}

	//There is no need to assign a struct value to a variable before creating a pointer, 
	//and the address operator can be used directly with the literal struct syntax
	{
		type Product struct {
			name, category string
			price float64
		}

		calcTax := func (product *Product) {
			if product.price > 100 {
				product.price += product.price * 0.2 
			}
		}

		kayak := &Product { "Kayak", "Watersports", 275 }

		calcTax(kayak)

		fmt.Println("Name:", kayak.name, "Category:", kayak.category, "Price:", kayak.price)
	}

	//Creating pointers directly from values can help make code more concise
	{
		type Product struct {
			name, category string
			price float64
		}

		calcTax := func (product *Product) *Product {
			if product.price > 100 {
				product.price += product.price * 0.2
			}
			return product
		}

		kayak := calcTax(&Product{
			name: "Kayak",
			category: "Watersports",
			price: 275,
		})

		fmt.Println("Name:", kayak.name, "Category:", kayak.category, "Price:", kayak.price)
	}

	//A constructor function is responsible for creating struct values using values received through parameters
	{
		type Product struct {
			name, category string
			price float64
		}

		newProduct := func (name, category string, price float64) *Product {
			return &Product{ name, category, price }
		}

		products := [2]*Product {
			newProduct( "Kayak", "Watersports", 275 ),
			newProduct( "Hat", "Skiing", 42.50 ),
		}

		for _, p := range products {
			fmt.Println("Name:", p.name, "Category:", p.category, "Price:", p.price)
		}
	}

	//Pointers can also be used for struct fields, including pointers to other struct types
	{
		type Supplier struct {
			name, city string
		}

		type Product struct {
			name, category string
			price float64
			*Supplier
		}

		newProduct := func(name, category string, price float64, supplier *Supplier) *Product {
			return &Product{ name, category, price, supplier }
		}

		acme := &Supplier{ "Acme Co", "New York" }

		products := [2]*Product {
			newProduct("Kayak", "Watersports", 275, acme),
			newProduct("Hat", "Skiing", 42.50, acme),
		}

		for _, p := range products {
			fmt.Println("Name:", p.name, "Supplier:", p.Supplier.name, p.Supplier.city)
		}
	}
}