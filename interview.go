package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// Customer struct represents the structure of customer data
type Customer struct {
	FirstName string
	LastName  string
	Email     string
	Gender    string
	IPAddress string
}

// readCustomers reads customer data from the given CSV file
func readCustomers(filename string) ([]Customer, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var customers []Customer
	for _, line := range lines {
		if len(line) != 5 {
			log.Printf("Skipping invalid line: %v", line)
			continue
		}
		customer := Customer{
			FirstName: line[0],
			LastName:  line[1],
			Email:     line[2],
			Gender:    line[3],
			IPAddress: line[4],
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

// getDomain returns the domain from an email address
func getDomain(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}

func main() {
	filename := "customers.csv"
	customers, err := readCustomers(filename)
	if err != nil {
		log.Fatalf("Error reading customers from file: %v", err)
	}

	domainCounts := make(map[string]int)
	for _, customer := range customers {
		domain := getDomain(customer.Email)
		if domain != "" {
			domainCounts[domain]++
		}
	}

	// Sorting the domains
	var domains []string
	for domain := range domainCounts {
		domains = append(domains, domain)
	}
	sort.Strings(domains)

	// Displaying the sorted list of domains along with the customer counts
	for _, domain := range domains {
		fmt.Printf("Domain: %s, Customers: %d\n", domain, domainCounts[domain])
	}
}
