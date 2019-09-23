package main

import (
	"fmt"
	"strconv"
	"strings"
)

// An WorldMap represents set of attributes that belongs to a map
type WorldMap struct {
	cities     map[string]*City
	aliens     map[uint]*Alien
	directions map[string]Direction
}

// NewWorldMap returns instance of WorldMap
func NewWorldMap() *WorldMap {
	return &WorldMap{
		cities:     make(map[string]*City),
		aliens:     make(map[uint]*Alien),
		directions: InitDirections(),
	}
}

// AddCities add cities to map and associates neighbours
func (w *WorldMap) AddCities(file string) error {
	lines := strings.Split(file, "\n")

	for _, line := range lines {
		words := strings.Split(line, " ")

		if len(words) > 1 {
			err := w.parseAndAddCity(TrimSpace(words[0]), words[1:])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// AddAliens adds aliens to random cities in the map
func (w *WorldMap) AddAliens(n uint) error {
	keys := make([]string, 0, len(w.cities))
	for k := range w.cities {
		keys = append(keys, k)
	}

	cities, err := Sample(keys, n)
	if err != nil {
		return err
	}

	for i, value := range cities {
		w.aliens[uint(i)] = NewAlien(strconv.Itoa(i), w.cities[value])
	}

	return nil
}

// Simulate the attack of aliens on world map
func (w *WorldMap) Simulate(steps int) {
	for iteration := 0; iteration < steps; iteration++ {
		fmt.Printf("Iteration %v\n", iteration+1)

		newCities := make(map[string]uint)

		if aliensLeft := w.anyAliensLeft(); !aliensLeft {
			fmt.Println("No Aliens Left!")
			fmt.Println(w)
			break
		}

		for counter := range w.aliens {
			if !w.aliens[counter].IsDestroyed() {
				oldCity := w.aliens[counter].city.name
				newCity := w.aliens[counter].Move()

				fmt.Printf("\nAlien %v moved from City %s to %s\n", counter, oldCity, newCity.name)

				if _, ok := newCities[newCity.name]; ok {
					newCities[newCity.name]++
				} else {
					newCities[newCity.name] = 1
				}
			}
		}

		for city := range newCities {
			if newCities[city] > 1 {
				w.cities[city].Destroy()
			}
		}

		fmt.Println(w)
	}
}

func (w *WorldMap) anyAliensLeft() bool {
	for counter := range w.aliens {
		if !w.aliens[counter].IsDestroyed() {
			return true
		}
	}

	return false
}

func (w *WorldMap) parseAndAddCity(cityName string, neighbours []string) error {
	city := w.createCity(cityName)

	for _, neighbourDirection := range neighbours {
		pairs := strings.Split(neighbourDirection, "=")

		directionName, neighbourName := TrimSpace(pairs[0]), TrimSpace(pairs[1])

		neighbour := w.createCity(neighbourName)
		direction, ok := w.directions[directionName]
		if !ok {
			return fmt.Errorf("Invalid direction %s", directionName)
		}

		err := city.SetNeighbour(direction, neighbour)
		if err != nil {
			return err
		}
	}

	return nil
}

func (w *WorldMap) createCity(cityName string) *City {
	if _, ok := w.cities[cityName]; !ok {
		w.cities[cityName] = NewCity(cityName)
	}

	return w.cities[cityName]
}

func (w *WorldMap) String() string {
	var cities []string
	for _, city := range w.cities {
		cities = append(cities, city.String())
	}

	var aliens []string
	for _, alien := range w.aliens {
		aliens = append(aliens, alien.String())
	}

	return strings.Join(append(cities, aliens...), "\n")
}
