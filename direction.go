package main

const (
	// EAST direction
	EAST string = "EAST"

	// WEST direction
	WEST string = "WEST"

	// NORTH direction
	NORTH string = "NORTH"

	// SOUTH direction
	SOUTH string = "SOUTH"
)

// Direction stores name of current direction and its inverse direction
type Direction struct {
	name    string
	inverse *Direction
}

// InitDirections creates a map containing all supported directions
func InitDirections() map[string]Direction {
	eastDirection := Direction{
		name: EAST,
	}
	westDirection := Direction{
		name: WEST,
	}
	northDirection := Direction{
		name: NORTH,
	}
	southDirection := Direction{
		name: SOUTH,
	}

	eastDirection.inverse = &westDirection
	westDirection.inverse = &eastDirection
	northDirection.inverse = &southDirection
	southDirection.inverse = &northDirection

	return map[string]Direction{
		EAST:  eastDirection,
		WEST:  westDirection,
		NORTH: northDirection,
		SOUTH: southDirection,
	}
}
