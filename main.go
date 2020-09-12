package main

import "fmt"

func main() {
	expressLine := NewBus("ExpressLine")
	expressLine.Add(Passenger{ssn: "001"})
	expressLine.Add(Passenger{ssn: "002"})

	// Get a list of ssns
	ssns := expressLine.List()
	fmt.Printf("This bus carries %d passengers, here are their SSN's: %v\n", len(ssns), ssns)
}
