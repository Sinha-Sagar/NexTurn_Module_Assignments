package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Product dataType
type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

type InventoryManager struct {
	products []Product
}

func NewInventoryManager() *InventoryManager {
	return &InventoryManager{
		products: make([]Product, 0),
	}
}

func (im *InventoryManager) AddProduct(id int, name string, priceStr string, stock int) error {
	for _, p := range im.products {
		if p.ID == id {
			return fmt.Errorf("product with ID %d already exists", id)
		}
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return errors.New("invalid price format")
	}

	if price <= 0 {
		return errors.New("price must be greater than zero")
	}
	if stock < 0 {
		return errors.New("stock cannot be negative")
	}
	if name == "" {
		return errors.New("product name cannot be empty")
	}

	product := Product{
		ID:    id,
		Name:  name,
		Price: price,
		Stock: stock,
	}

	im.products = append(im.products, product)
	return nil
}

func (im *InventoryManager) UpdateStock(id int, newStock int) error {
	if newStock < 0 {
		return errors.New("stock cannot be negative")
	}

	for i := range im.products {
		if im.products[i].ID == id {
			im.products[i].Stock = newStock
			return nil
		}
	}

	return fmt.Errorf("product with ID %d not found", id)
}

func (im *InventoryManager) SearchByID(id int) (*Product, error) {
	for i := range im.products {
		if im.products[i].ID == id {
			return &im.products[i], nil
		}
	}
	return nil, fmt.Errorf("product with ID %d not found", id)
}

func (im *InventoryManager) SearchByName(name string) []Product {
	var results []Product
	searchTerm := strings.ToLower(name)

	for _, p := range im.products {
		if strings.Contains(strings.ToLower(p.Name), searchTerm) {
			results = append(results, p)
		}
	}

	return results
}

func (im *InventoryManager) SortByPrice() {
	sort.Slice(im.products, func(i, j int) bool {
		return im.products[i].Price < im.products[j].Price
	})
}

func (im *InventoryManager) SortByStock() {
	sort.Slice(im.products, func(i, j int) bool {
		return im.products[i].Stock < im.products[j].Stock
	})
}

func (im *InventoryManager) DisplayInventory() {
	if len(im.products) == 0 {
		fmt.Println("Inventory is empty")
		return
	}

	fmt.Printf("\n%-5s | %-20s | %-10s | %-8s\n", "ID", "Name", "Price ($)", "Stock")
	fmt.Println(strings.Repeat("-", 50))

	for _, p := range im.products {
		fmt.Printf("%-5d | %-20s | %10.2f | %-8d\n",
			p.ID, p.Name, p.Price, p.Stock)
	}
	fmt.Println()
}

func main() {
	inventory := NewInventoryManager()
	samples := []struct {
		id    int
		name  string
		price string
		stock int
	}{
		{1, "Laptop", "999.99", 10},
		{2, "Mouse", "29.99", 50},
		{3, "Keyboard", "59.99", 30},
		{4, "Monitor", "299.99", 15},
		{5, "USB Cable", "9.99", 100},
	}

	for _, s := range samples {
		err := inventory.AddProduct(s.id, s.name, s.price, s.stock)
		if err != nil {
			fmt.Printf("Error adding %s: %v\n", s.name, err)
		}
	}

	fmt.Println("Original Inventory:")
	inventory.DisplayInventory()

	err := inventory.UpdateStock(1, 15)
	if err != nil {
		fmt.Printf("Error updating stock: %v\n", err)
	}

	product, err := inventory.SearchByID(1)
	if err != nil {
		fmt.Printf("Search error: %v\n", err)
	} else {
		fmt.Printf("\nFound product by ID: %+v\n", *product)
	}

	results := inventory.SearchByName("key")
	fmt.Printf("\nProducts containing 'key':\n")
	for _, p := range results {
		fmt.Printf("%+v\n", p)
	}

	fmt.Println("\nInventory sorted by price:")
	inventory.SortByPrice()
	inventory.DisplayInventory()

	fmt.Println("\nInventory sorted by stock:")
	inventory.SortByStock()
	inventory.DisplayInventory()
}
