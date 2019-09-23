package main

import "fmt"

// Alien stores properties like name. city and is alien destroyed
type Alien struct {
	name      string
	city      *City
	destroyed bool
}

// NewAlien returns new alien instance
func NewAlien(name string, city *City) *Alien {
	alien := &Alien{
		name:      name,
		city:      city,
		destroyed: false,
	}

	alien.city.AddResident(alien)
	return alien
}

// Move alien to neighbour city
func (a *Alien) Move() *City {
	city, err := a.city.GetNeighbour()
	if err != nil {
		fmt.Printf("\nAlien %s is trapped in city %s!\n", a.name, a.city.name)
		return a.city
	}

	oldCity := a.city
	a.city = city
	a.city.AddResident(a)
	oldCity.DeleteResident(a)

	return a.city
}

// Destroy alien
func (a *Alien) Destroy() {
	a.destroyed = true
}

// IsDestroyed check if alien is destroyed or not
func (a *Alien) IsDestroyed() bool {
	return a.destroyed
}

// String print alien information in string format
func (a *Alien) String() string {
	state := "'Alive and is in city " + a.city.name
	if a.destroyed {
		state = "Dead!"
	}

	return fmt.Sprintf("Alien %s is %s", a.name, state)
}
