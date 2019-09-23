package main

import (
	"errors"
	"fmt"
	"strings"
)

// An City represents set of attributes that belongs to a city
type City struct {
	name       string
	neighbours map[Direction]*City
	residents  map[string]*Alien
	destroyed  bool
}

// NewCity returns object of city
func NewCity(name string) *City {
	return &City{
		name:       name,
		neighbours: make(map[Direction]*City),
		residents:  make(map[string]*Alien),
	}
}

// SetNeighbour function sets neighbour of current city to provided city
func (c *City) SetNeighbour(direction Direction, city *City) error {
	if city == c {
		return fmt.Errorf("City `%s` cannot be the neighbour of itself", c.name)
	}

	for direction, neighbour := range c.neighbours {
		if city == neighbour {
			if _, ok := c.neighbours[direction]; ok && c.neighbours[direction] == city {
				return nil
			}

			return fmt.Errorf("City `%s` is already a neighbour to city `%s`", city.name, c.name)
		}
	}

	c.neighbours[direction] = city
	city.neighbours[*direction.inverse] = c

	return nil
}

// GetNeighbour returns a neighbour of current city
func (c *City) GetNeighbour() (*City, error) {
	// [TODO] Add random logic
	for _, value := range c.neighbours {
		return value, nil
	}

	return nil, errors.New("No neighbour found")
}

// AddResident adds a new alien as resident of current city
func (c *City) AddResident(alien *Alien) {
	if _, ok := c.residents[alien.name]; ok {
		return
	}

	c.residents[alien.name] = alien
}

// DeleteResident deletes resident (alien) from current city
func (c *City) DeleteResident(alien *Alien) {
	delete(c.residents, alien.name)
}

// Destroy function completely destroy current city and removes its reference from its neighbours
func (c *City) Destroy() {
	for direction, city := range c.neighbours {
		delete(city.neighbours, *direction.inverse)
	}

	c.neighbours = make(map[Direction]*City)
	c.destroyed = true

	names := []string{}
	for name := range c.residents {
		c.residents[name].Destroy()
		names = append(names, name)
	}

	fmt.Printf("\n%s has been destroyed by %s\n", c.name, strings.Join(names[:], ", "))
}

func (c *City) String() string {
	if c.destroyed {
		return fmt.Sprintf("City %s is destroyed!", c.name)
	}

	return fmt.Sprintf("City %s %s", c.name, c.neighbourString())
}

func (c *City) neighbourString() string {
	var neighbours []string

	for direction := range c.neighbours {
		name := "None"
		if _, ok := c.neighbours[direction]; ok {
			name = c.neighbours[direction].name
		}

		neighbours = append(neighbours, fmt.Sprintf("%v=%v", direction.name, name))
	}

	return strings.Join(neighbours, " ")
}
