package wrappers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Beverage represents the structure of a beverage
type Beverage struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// BaseURL is the base endpoint for the beverage API
const BaseURL = "http://localhost:5000/beverages"

// CreateBeverage sends a POST request to create a new beverage
func CreateBeverage(name, beverageType string) (*Beverage, error) {
	beverage := Beverage{Name: name, Type: beverageType}
	data, err := json.Marshal(beverage)
	if err != nil {
		return nil, fmt.Errorf("error marshalling data: %w", err)
	}

	resp, err := http.Post(BaseURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("error creating beverage: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to create beverage: %s", string(body))
	}

	var createdBeverage Beverage
	err = json.NewDecoder(resp.Body).Decode(&createdBeverage)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &createdBeverage, nil
}

// GetAllBeverages sends a GET request to retrieve all beverages
func GetAllBeverages() ([]Beverage, error) {
	resp, err := http.Get(BaseURL)
	if err != nil {
		return nil, fmt.Errorf("error fetching beverages: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to fetch beverages: %s", string(body))
	}

	var beverages []Beverage
	err = json.NewDecoder(resp.Body).Decode(&beverages)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return beverages, nil
}

// GetBeverage sends a GET request to retrieve a beverage by ID
func GetBeverage(id int) (*Beverage, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%d", BaseURL, id))
	if err != nil {
		return nil, fmt.Errorf("error fetching beverage: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to fetch beverage: %s", string(body))
	}

	var beverage Beverage
	err = json.NewDecoder(resp.Body).Decode(&beverage)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &beverage, nil
}

// UpdateBeverage sends a PUT request to update a beverage by ID
func UpdateBeverage(id int, name, beverageType string) (*Beverage, error) {
	beverage := Beverage{Name: name, Type: beverageType}
	data, err := json.Marshal(beverage)
	if err != nil {
		return nil, fmt.Errorf("error marshalling data: %w", err)
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/%d", BaseURL, id), bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error updating beverage: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to update beverage: %s", string(body))
	}

	var updatedBeverage Beverage
	err = json.NewDecoder(resp.Body).Decode(&updatedBeverage)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &updatedBeverage, nil
}

// DeleteBeverage sends a DELETE request to remove a beverage by ID
func DeleteBeverage(id int) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%d", BaseURL, id), nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error deleting beverage: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("failed to delete beverage: %s", string(body))
	}

	return nil
}
