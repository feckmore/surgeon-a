package surgeon

import "time"

// ID -trying this out
type ID int

// Surgeon is the core business entity... along with the methods defined in service.go,
// it makes up the domain logic
type Surgeon struct {
	ID ID

	Prefix    string
	Firstname string
	Lastname  string
	Suffix    string

	Organization string
	Email        string
	Website      string

	Address1 string
	Address2 string
	City     string
	State    string
	Zip      string
	Country  string

	Phone string
	Fax   string
	Photo string

	Latitude  float64
	Longitude float64

	Created time.Time
	Updated time.Time
}
