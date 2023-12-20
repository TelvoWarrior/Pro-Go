package main

import ( 
	"fmt"
	"composition/store"
)
func main() {
	fmt.Println("Hello, Composition!")

	kayak := store.NewProduct("Kayak", "Watersports", 275)
	lifejacket := &store.Product{ Name: "Lifejacket", Category: "Watersports"}

	for _, p := range []*store.Product { kayak, lifejacket } {
		fmt.Println("Name:", p.Name, "Category:", p.Category, "Price:", p.Price(0.2))
	}

	boats := []*store.Boat {
		store.NewBoat("Kayak", 275, 1, false),
		store.NewBoat("Canoe", 400, 3, false),
		store.NewBoat("Tender", 650.25, 2, true),
	}

	for _, b := range boats {
		fmt.Println("Conventional:", b.Product.Name, "Direct:", b.Name)
		fmt.Println("Boat:", b.Name, "Price:", b.Price(0.2))
	}

	rentals := []*store.RentalBoat {
		store.NewRentalBoat("Rubber Ring", 10, 1, false, false, "N/A", "N/A"),
		store.NewRentalBoat("Yacht", 5000, 5, true, true, "Bob", "Alice"),
		store.NewRentalBoat("Super Yacht", 100000, 15, true, true, "Dora", "Charlie"),
	}

	for _, r := range rentals {
		fmt.Println("Rental Boat:", r.Name, "Rental Price:", r.Price(0.2), "Captain:", r.Captain)
	}

	product := store.NewProduct("Kayak", "Watersports", 279)

	deal := store.NewSpecialDeal("Weekend Special", product, 50)

	Name, price, Price := deal.GetDetails()

	fmt.Println("Name:", Name)
	fmt.Println("Price field:", price)
	fmt.Println("Price method:", Price)
}