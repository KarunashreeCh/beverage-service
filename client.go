package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Beverage represents the structure of a beverage
type Beverage struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Cost     float64 `json:"cost,omitempty"`
}

const baseURL = "http://localhost:5000/beverages"

// CreateBeverage sends a POST request to create a new beverage
func CreateBeverage(beverage Beverage) {
	data, err := json.Marshal(beverage)
	if err != nil {
		log.Fatalf("Error marshalling data: %v", err)
	}

	resp, err := http.Post(baseURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Error creating beverage: %v", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Create Response: %s\n", body)
}

// GetAllBeverages sends a GET request to fetch all beverages
func GetAllBeverages() {
	resp, err := http.Get(baseURL)
	if err != nil {
		log.Fatalf("Error fetching beverages: %v", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("All Beverages: %s\n", body)
}

// GetBeverage sends a GET request to fetch a single beverage by ID
func GetBeverage(id int) {
	resp, err := http.Get(fmt.Sprintf("%s/%d", baseURL, id))
	if err != nil {
		log.Fatalf("Error fetching beverage: %v", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Beverage %d: %s\n", id, body)
}

// UpdateBeverage sends a PUT request to update a beverage by ID
func UpdateBeverage(id int, beverage Beverage) {
	data, err := json.Marshal(beverage)
	if err != nil {
		log.Fatalf("Error marshalling data: %v", err)
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/%d", baseURL, id), bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error updating beverage: %v", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Update Response: %s\n", body)
}

// DeleteBeverage sends a DELETE request to remove a beverage by ID
func DeleteBeverage(id int) {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%d", baseURL, id), nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error deleting beverage: %v", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Delete Response: %s\n", body)
}

func main() {
	// Example usage
	fmt.Println("Creating a beverage...")
	CreateBeverage(Beverage{Name: "Coke", Quantity: 10, Cost: 1.5})

	fmt.Println("\nGetting all beverages...")
	GetAllBeverages()

	fmt.Println("\nGetting a single beverage...")
	GetBeverage(1)

	fmt.Println("\nUpdating a beverage...")
	UpdateBeverage(1, Beverage{Name: "Pepsi", Quantity: 20})

	fmt.Println("\nDeleting a beverage...")
	DeleteBeverage(1)
}
