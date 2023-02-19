package store

type RentalBoat struct {
	*Boat
	IncludeCrew bool
	*Crew
}

/* The RentalBoat type is composed using the *Boat type, which is, in turn,
* composed using the *Product type, forming a chain. Go performs promotion so
* that the fields defined by all three types in the chain can be accessed
* directly, as shown in Listing 13-11.  */
func NewRentalBoat(name string, price float64, capacity int, motorized,
	crewed bool, captain, firstOfficer string) *RentalBoat {
	return &RentalBoat{NewBoat(name, price, capacity, motorized), crewed,
		&Crew{captain, firstOfficer}}
}

type Crew struct{ Captain, FirstOfficer string }
