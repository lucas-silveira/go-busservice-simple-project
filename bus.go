package main

import "fmt"

// Bus carries Passengers from A to B if they have a valid bus ticket.
type Bus struct {
	name       string
	passengers Passengers
}

// NewBus returns a new Bus with an empty passenger set.
func NewBus(name string) Bus {
	bus := Bus{}
	bus.name = name
	bus.passengers = NewPassengerSet()
	return bus
}

// Add adds a single passenger to the Bus. For brevity, we don't care too much about accidentally adding the same Passenger more than once.
func (b *Bus) Add(p Passenger) {
	b.passengers.Add(p)
	fmt.Printf("%s: boarded passenger with SSN %q\n", b.name, p.ssn)
}

// Remove removes a single Passenger from the Bus.
func (b *Bus) Remove(p Passenger) {
	b.passengers.Remove(p)
	fmt.Printf("%s: unboarded passenger with SSN %q\n", b.name, p.ssn)
}

// List asks Passengers for a SSN list and returns it.
func (b Bus) List() []string {
	return b.passengers.List()
}

// VisitPassengers calls function visitor for each Passenger on the bus.
func (b *Bus) VisitPassengers(visitor func(Passenger)) {
	b.passengers.Visit(visitor)
}

// FindPassenger returns the Passenger that matches the given SSN, if found. Otherwise, an empty Passenger is returned.
func (b *Bus) FindPassenger(ssn string) Passenger {
	if p, ok := b.passengers[ssn]; ok {
		return p
	}
	return Passenger{} // A nobody.
}

// UpdatePassengers calls function visitor for each Passenger on the bus. Passengers are passed by reference and may be modified.
func (b *Bus) UpdatePassengers(visitor func(*Passenger)) {
	ps := make(map[string]Passenger, len(b.passengers))
	for ssn, p := range b.passengers {
		visitor(&p)
		ps[ssn] = p
	}
	b.passengers = ps
}
