# Alien Invasion

## Problem Statement

Mad aliens are about to invade the earth and you are tasked with simulating the
invasion.

You are given a map containing the names of cities in the non-existent world of
X. The map is in a file, with one city per line. The city name is first,
followed by 1-4 directions (NORTH, SOUTH, EAST, or WEST). Each one represents a
road to another city that lies in that direction.

For example:
Foo NORTH=Bar WEST=Baz SOUTH=Qu-ux
Bar SOUTH=Foo WEST=Bee

The city and each of the pairs are separated by a single space, and the
directions are separated from their respective cities with an equals (=) sign.
You should create N aliens, where N is specified as a command-line argument.
These aliens start out at random places on the map, and wander around randomly,
following links. Each iteration, the aliens can travel in any of the directions
leading out of a city. In our example above, an alien that starts at Foo can go
NORTH to Bar, WEST to Baz, or SOUTH to Qu-ux.
When two aliens end up in the same place, they fight, and in the process kill
each other and destroy the city. When a city is destroyed, it is removed from
the map, and so are any roads that lead into or out of it.
In our example above, if Bar were destroyed the map would now be something
like:
Foo WEST=Baz SOUTH=Qu-ux
Once a city is destroyed, aliens can no longer travel to or through it. This
may lead to aliens getting "trapped".
You should create a program that reads in the world map, creates N aliens, and
unleashes them. The program should run until all the aliens have been
destroyed, or each alien has moved at least 10,000 times. When two aliens
fight, print out a message like:

Bar has been destroyed by alien 10 and alien 34!

## Assumption

- Directions would be provided in caps: NORTH, SOUTH, EAST, WEST


## Solution

To test this solution:

1. Update input.txt file to represent your map. for reference, it contains basic map:

```
Foo NORTH=Bar WEST=Baz SOUTH=Qu-ux
Bar SOUTH=Foo WEST=Bee
```

2. Build Project:

```
$ go build -n invasion
```

3. Run Project and Enter number of aliens:

```
$ ./invasion

Number of Aliens :
5
```

