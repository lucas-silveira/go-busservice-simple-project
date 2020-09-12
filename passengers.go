package main

// Passenger represents a bus passenger, uniquely identified by their SSN.
type Passenger struct {
	ssn        string
	seatNumber uint8
}

// Passengers represents a set of Passengers, using their SSN as key.
type Passengers map[string]Passenger

// NewPassengerSet returns an empty set of Passengers, ready to use.
func NewPassengerSet() Passengers {
	return make(map[string]Passenger)
}

// Add adds a Passenger to Passengers. The Passenger will be overwritten if exists.
func (ps Passengers) Add(p Passenger) {
	ps[p.ssn] = p
}

// Remove removes a Passenger from Passengers.
func (ps Passengers) Remove(p Passenger) {
	delete(ps, p.ssn)
}

// Visit calls visitor once for every Passenger in the set.
func (ps Passengers) Visit(visitor func(Passenger)) {
	for _, one := range ps {
		visitor(one)
	}
}

// Find returns the Passenger with the given SSN. If none was found, an empty Passenger is returned.
func (ps Passengers) Find(ssn string) Passenger {
	if one, ok := ps[ssn]; ok {
		return one
	}
	return Passenger{}
}

// VisitUpdate calls visitor for each Passenger in the set. Updating their SSN's is not recommended.
func (ps *Passengers) VisitUpdate(visitor func(p *Passenger)) {
	for ssn, pp := range *ps {
		visitor(&pp)
		(*ps)[ssn] = pp
	}
}

// List returns the SSN's of all Passengers in the set.
func (ps Passengers) List() []string {
	ssns := make([]string, 0, len(ps))
	ps.Visit(func(p Passenger) { ssns = append(ssns, p.ssn) })
	return ssns
}
