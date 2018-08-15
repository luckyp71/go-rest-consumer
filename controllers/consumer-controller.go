package Controllers

import (
	"encoding/json"
	"fmt"
	m "go-rest-consumer/models"
	"log"
	"net/http"
)

var categories []m.Category

func Controllers() {

	// Sample REST URI with HTTP Method: GET
	url := "http://localhost:8090/menu/categories"

	// Invoke the RESTful web service to get category list
	getCategories(url)

}

func getCategories(url string) {
	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Create a client.
	client := &http.Client{}

	// Send the request via a client and do sends an HTTP request and returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Callers should close the response body when done reading from it
	defer resp.Body.Close()

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&categories); err != nil {
		log.Println(err)
	}

	for _, category := range categories {
		fmt.Println("Category Name: ", category.Category)
	}
}
